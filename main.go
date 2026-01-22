package main

import (
	"fmt"
	"os"
	logger "pgadminserverautofill/main/src"
	"strings"
	"text/template"
)

func main() {

	requiredEnvVars := []string{
		"LOG_LEVEL",
	}

	for _, key := range requiredEnvVars {
		checkEnvExists(key)
	}

	log := logger.New()

	log.Debug("Listing all environment variables:")
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		log.Debug("environment variable", "key", pair[0], "value", pair[1])
	}

	log.Debug("-------------------------------------------------------")

	requiredEnvVars = []string{
		"LOG_LEVEL",
		"PG_HOST",
		"PG_PORT",
	}

	for _, key := range requiredEnvVars {
		checkEnvExists(key)
	}

	data := make(map[string]string)
	for _, e := range os.Environ() {
		parts := strings.SplitN(e, "=", 2)
		data[parts[0]] = parts[1]
	}

	tmpl, err := template.ParseFiles("servers.tpl")
	if err != nil {
		panic(err)
	}

	err = os.MkdirAll("/output", 0o755)
	if err != nil {
		panic(err)
	}

	out, err := os.Create("/output/servers.json")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	err = tmpl.Execute(out, data)
	if err != nil {
		panic(err)
	}

	d, err := os.ReadFile("/output/servers.json")
	if err != nil {
		panic(err)
	}

	log.Debug("File content:", "content", string(d))
}

func checkEnvExists(key string) {
	if _, exists := os.LookupEnv(key); !exists {
		panic(fmt.Sprintf("Environment variable %s is required but not set", key))
	}
}
