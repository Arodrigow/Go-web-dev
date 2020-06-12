package main

import (
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func register(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	if req.Method == http.MethodPost {
		fName := req.FormValue("fName")
		lName := req.FormValue("lName")
		ID := req.FormValue("id")
		password := req.FormValue("password")

		if userExist(ID) {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}

		bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
		handleError(err)

		query := fmt.Sprintf(`INSERT INTO users (fName, lName, userid, password) VALUES ("%s", "%s", "%s", "%s");`, fName, lName, ID, string(bs))
		stmt, err := db.Prepare(query)
		handleError(err)
		defer stmt.Close()

		_, err = stmt.Exec()
		handleError(err)

		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "register.gohtml", nil)
}
