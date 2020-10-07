package main

func main() {
	a := App{}

	// Inicialize with the environment variables
	a.Initialize(
		DefaultEnv("APP_DB_USERNAME", "postgres"),
		DefaultEnv("APP_DB_PASSWORD", "postgres"),
		DefaultEnv("APP_DB_NAME", "postgres"),
	)

	a.Run(":8010")
}
