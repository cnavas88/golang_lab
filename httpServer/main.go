package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Process dynamic request
	http.HandleFunc("/", welcome)
	http.HandleFunc("/up", upWebsite)

	// Serving static files
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Accept connections
	http.ListenAndServe(":8080", nil)
}

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my website!")
}

func upWebsite(w http.ResponseWriter, r *http.Request) {
	website := "http://google.com"

	_, err := http.Get(website)
	if err != nil {
		fmt.Errorf("google not response", http.StatusNotFound)
	}

	fmt.Fprintf(w, "Google is up!")
}
