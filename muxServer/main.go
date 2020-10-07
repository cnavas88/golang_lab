package main

import (
	"fmt"
)

func main() {
	a := App{}

	psqlInfo := map[string]string{
		"host":     DefaultEnv("APP_DB_HOST", "localhost"),
		"port":     DefaultEnv("APP_DB_PORT", "5432"),
		"username": DefaultEnv("APP_DB_USERNAME", "postgres"),
		"password": DefaultEnv("APP_DB_PASSWORD", ""),
		"dbname":   DefaultEnv("APP_DB_NAME", "postgres"),
	}

	fmt.Println(psqlInfo)

	// Inicialize with the environment variables
	a.Initialize(psqlInfo)

	a.Run(":8010")
}
