package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

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

// Route Handlers
func (a *App) getProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// Atoi is equivalent to ParseInt
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	p := product{ID: id}
	if err := p.getProduct(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithJSON(w, http.StatusNotFound, "Product Not found")
		default:
			respondWithJSON(w, http.StatusInternalServerError, err.Error())
		}
	}

	respondWithJSON(w, http.StatusOK, p)
}

func (a *App) getProducts(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	count, start = controlValues(count, start)

	products, err := getProducts(a.DB, start, count)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, products)
}

func controlValues(count, start int) (int, int) {
	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	return count, start
}
