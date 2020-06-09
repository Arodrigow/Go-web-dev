package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)
	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("css"))))
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	handleError(w, err)
}

func register(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "register.gohtml", nil)
	handleError(w, err)
}

func login(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "login.gohtml", nil)
	handleError(w, err)
}

func handleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
