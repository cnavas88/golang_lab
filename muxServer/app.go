package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// App struct exposes references to the router and the database
// that the application uses. To be useful and testable, will need
// two methods that initialize and run applicacion
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

// Initialize have the details required to connect to the database
func (a *App) Initialize() {
	// connectionString :=
	// 	fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	// 		p["host"],
	// 		p["port"],
	// 		p["user"],
	// 		p["password"],
	// 		p["dbname"],
	// 	)

	connectionString := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal(err)
	}

	// Only initialize when we've connected to database
	a.Router = mux.NewRouter()
}

// Run method will simply start the app.
func (a *App) Run(addr string) {}
