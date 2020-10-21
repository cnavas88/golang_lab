package main

import (
	"fmt"
	"net/http"
	"text/template"
)

// ContactDetails store the basic details to send an email
type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

func main() {
	tmpl := template.Must(template.ParseFiles("forms.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		details := ContactDetails{
			Email:   r.FormValue("email"),
			Subject: r.FormValue("subject"),
			Message: r.FormValue("message"),
		}

		// do something with details
		fmt.Println(details)

		_ = details

		tmpl.Execute(w, struct{ Success bool }{true})
	})

	http.ListenAndServe(":8000", nil)
}
