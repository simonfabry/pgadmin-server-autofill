package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

func main() {

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0], pair[1])
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

	out, err := os.Create("servers.json")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	err = tmpl.Execute(out, data)
	if err != nil {
		panic(err)
	}
}
