package main

import "net/http"

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8080", nil)

	//Other option
	//log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
}
