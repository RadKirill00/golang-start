package main

import (
	"html/template"
	"net/http"
)

type User struct {
	Name    string
	Age     int
	Gender  string
	Hobbies []string
}

func main() {
	user := User{"Kirill", 17, "Man", []string{"Football", "Reading the books", "Play on PS4"}}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("golang-start/index.html")
		tmpl.Execute(w, user)
	})
	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		tmpl2, _ := template.ParseFiles("golang-start/contact.html")
		tmpl2.Execute(w, nil)
	})
	http.ListenAndServe(":8080", nil)
}
