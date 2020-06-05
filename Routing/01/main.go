package main

import (
	"io"
	"net/http"
)

func def(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "This is the default page")
}

func d(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "This is a page named dog")
}

func m(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "This is a page named me")
}
func main() {
	http.HandleFunc("/", def)
	http.HandleFunc("/dog/", d)
	http.HandleFunc("/me/", m)

	http.ListenAndServe(":8080", nil)
}
