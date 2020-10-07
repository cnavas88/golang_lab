package main

/*
	To be able initialize the databse you have to generate the postgres container
	with docker. If you don't have docker you can download it by accessing:
	- https://docs.docker.com/get-docker/

	to create and inicialize the postgres container, execute the next command:
	# docker run -e POSTGRES_HOST_AUTH_METHOD=trust -it -p 5432:5432 -d postgres
*/

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
