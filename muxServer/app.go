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

// Initialize have the details required to connect to the database
func (a *App) Initialize(p map[string]string) {

	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		p["host"],
		p["port"],
		p["username"],
		p["password"],
		p["dbname"],
	)

	fmt.Println(connectionString)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)

	if err != nil {
		// panic error and return, we can't init application if don't
		// have connection to the database
		log.Fatal(err)
	}

	// Only initialize when we've connected to database
	a.Router = mux.NewRouter()
}

// Run method will simply start the app.
func (a *App) Run(addr string) {}
