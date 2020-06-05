package main

import (
	"io"
	"log"
	"net/http"
	"text/template"
)

func defa(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "This is the default page")
}

func d(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "This is a page named dog")
}

func m(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "tpl.gohtml", "This page is mine")
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))

}

func main() {
	http.HandleFunc("/", http.HandlerFunc(defa))
	http.HandleFunc("/dog/", http.HandlerFunc(d))
	http.HandleFunc("/me/", http.HandlerFunc(m))

	http.ListenAndServe(":8080", nil)
}
