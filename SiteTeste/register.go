package main

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"

	uuid "github.com/satori/go.uuid"
)

func register(w http.ResponseWriter, req *http.Request) {

	//if alreadyLoggedIn(req) {
	//	http.Redirect(w, req, "/", http.StatusSeeOther)
	//	return
	//}

	//Process post request
	if req.Method == http.MethodPost {

		//getting data from the webform
		Fname := req.FormValue("fName")
		Lname := req.FormValue("lName")
		ID := req.FormValue("id")
		Password := req.FormValue("password")

		//checking if the user already exists
		if userExist(ID) {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}

		//create session
		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
		addSession(ID, c.Value)

		//Storing users in mySQL
		bs, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		addUser(ID, Fname, Lname, string(bs))
		// redirect
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	err := tpl.ExecuteTemplate(w, "register.gohtml", nil)
	handleError(w, err)
}
