package main

import (
	"log"
	"os"
	"testing"
)

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS products
(
	id SERIAL,
	name TEXT NOT NULL,
	price NUMERIC(10,2) NOT NULL DEFAULT 0.00,
	CONSTRAINT products_pkey PRIMERY KEY (id)
)
`

var a = App{}

func TestMain(m *testing.M) {
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
	)

	ensureTableExists()

	// All the tests are executed by calling m.Run()
	code := m.Run()

	// after run test clean the database up
	clearTable()
	os.Exit(code)
}

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.DB.Exec("DELETE FROM products")
	a.DB.Exec("ALTER SEQUENCE products_id_seq RESTART WITH 1")
}

// func TestEmptyTable(t *testing.T) {
// 	clearTable()

// 	req, _ := http.NewRequest("GET", "/products", nil)
// 	response := executeRequest(req)

// 	checkResponseCode(t, http.StatusOK, response.Code)

// 	if body := response.Body.String(); body != "[]" {
// 		t.Errorf("Expected an empty array. Got %s", body)
// 	}
// }

// func executeRequest(req *http.Request) *httptest.ResponseRecorder {
// 	rr := httptest.NewRecorder()
// 	a.Router.ServeHTTP(rr, req)

// 	return rr
// }

// func checkResponseCode(t *testing.T, expected int, actual int) {
// 	if expected != actual {
// 		t.Errorf("Expected response code %d. Got %d\n", expeceted, actual)
// 	}
// }