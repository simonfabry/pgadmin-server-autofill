{
    "Servers": {
        "1": {
            "Name": "Minimally Defined Server",
            "Group": "Default",
            "Port": {{ .PG_PORT }},
            "Username": "postgres",
            "Host": "{{ .PG_HOST }}",
            "SSLMode": "prefer",
            "MaintenanceDB": "postgres"
        }
    }
}