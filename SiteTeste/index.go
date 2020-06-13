package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, req *http.Request) {
	users := getAllUsers()
	if alreadyLoggedIn(req) {
		data := struct {
			Us       []user
			LoggedIn bool
		}{
			Us:       users,
			LoggedIn: true,
		}
		fmt.Println(data.LoggedIn)
		tpl.ExecuteTemplate(w, "index.gohtml", data)
		return
	}
	data := struct {
		Us       []user
		LoggedIn bool
	}{
		Us:       users,
		LoggedIn: false,
	}
	tpl.ExecuteTemplate(w, "index.gohtml", data)

}
