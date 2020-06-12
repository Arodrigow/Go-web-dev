package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, req *http.Request) {
	users := getAllUsers()
	if alreadyLoggedIn(req) {
		data := struct {
			us       []user
			loggedIn bool
		}{
			us:       users,
			loggedIn: true,
		}
		fmt.Println(data.loggedIn)
		tpl.ExecuteTemplate(w, "index.gohtml", data)
		return
	}
	data := struct {
		us       []user
		loggedIn bool
	}{
		us:       users,
		loggedIn: false,
	}
	tpl.ExecuteTemplate(w, "index.gohtml", data)

}
