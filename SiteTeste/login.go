package main

import (
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func login(w http.ResponseWriter, req *http.Request) {

	if alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	if req.Method == http.MethodPost {
		un := req.FormValue("id")
		p := req.FormValue("password")

		if tryLogIn(un, p) == false {
			http.Error(w, "Username or password do no match", http.StatusNotAcceptable)
			return
		}

		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
		newSession(un, c.Value)

		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "login.gohtml", nil)
}
