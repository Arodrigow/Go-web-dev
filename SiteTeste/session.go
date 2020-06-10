package main

import (
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func getUser(w http.ResponseWriter, req *http.Request) user {

	//First, create or find the cookie of the session
	c, err := req.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
	}
	http.SetCookie(w, c)
	var u user
	/*	//If theres a user linked to the cookie, get the user

		if un, ok := PROBLEMA; ok {
			u = PROBLEMA
		}
	*/
	return u
}
