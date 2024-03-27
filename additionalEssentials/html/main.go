package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type User struct {
	Name  string
	Email string
}

func main() {

	templates = template.Must(template.ParseGlob("*.html"))

	u := User{"caio", "caio@gmail.com"}

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		templates.ExecuteTemplate(w, "home.html", u)
	})

	fmt.Println("Porta 5000 aberta")

	log.Fatal(http.ListenAndServe(":5000", nil))
}
