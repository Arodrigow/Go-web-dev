package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func main() {
	http.HandleFunc("/", surf)
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))
	//This workds as well
	//http.Handle("/public/", http.FileServer(http.Dir(".")))

	http.ListenAndServe(":8080", nil)
}

func surf(w http.ResponseWriter, req *http.Request) {
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}
}
