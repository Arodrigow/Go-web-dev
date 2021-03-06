package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", set)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, req *http.Request) {

	cookie, err := req.Cookie("n-visits")
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "n-visits",
			Value: "0",
		}
	}

	count, err := strconv.Atoi(cookie.Value)
	if err != nil {
		log.Fatalln(err)
	}

	count++
	cookie.Value = strconv.Itoa(count)

	http.SetCookie(w, cookie)

	io.WriteString(w, cookie.Value)
}
