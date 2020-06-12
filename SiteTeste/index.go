package main

import "net/http"

func index(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) {
		c, _ := req.Cookie("session")
		un := getIDFromValue(c.Value)
		us := getUser(un)

		tpl.ExecuteTemplate(w, "index.gohtml", us)
		return
	}
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
