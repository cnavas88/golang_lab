package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// User new struct to example
type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

func main() {
	http.HandleFunc("/decode", decode)
	http.HandleFunc("/encode", encode)

	http.ListenAndServe(":8080", nil)
}

func decode(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)

	fmt.Fprintf(w, "%s %s is %d years old!", user.Firstname, user.Lastname, user.Age)
}

func encode(w http.ResponseWriter, r *http.Request) {
	carlos := User{
		Firstname: "Carlos",
		Lastname:  "Navas",
		Age:       31,
	}

	json.NewEncoder(w).Encode(carlos)
}
