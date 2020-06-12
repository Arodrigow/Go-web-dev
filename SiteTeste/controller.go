package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

//User's func

func userExist(id string) bool {
	var us user
	var iddb int
	query := fmt.Sprintf(`SELECT * FROM users WHERE userid="%s";`, id)
	err = db.QueryRow(query).Scan(&iddb, &us.FName, &us.LName, &us.ID, &us.Password)
	if err != nil {
		return false
	}
	return true
}

func getUserFromCookie(cookieValue string) bool {
	var iddb int
	var user user

	var userid string
	query := fmt.Sprintf(`SELECT userid FROM dbsession WHERE uuid="%s";`, cookieValue)
	err = db.QueryRow(query).Scan(&userid)
	if err != nil {
		return false
	}

	query = fmt.Sprintf(`SELECT * FROM users WHERE userid="%s";`, userid)
	err = db.QueryRow(query).Scan(&iddb, &user.FName, &user.LName, &user.ID, &user.Password)
	if err != nil {
		return false
	}

	return true
}

func getUser(userid string) user {
	var iddb int
	var user user

	query := fmt.Sprintf(`SELECT * FROM users WHERE userid="%s";`, userid)
	err = db.QueryRow(query).Scan(&iddb, &user.FName, &user.LName, &user.ID, &user.Password)
	handleError(err)

	return user
}

func getAllUsers() []user {
	var u user
	var users []user
	var iddb int

	query := fmt.Sprintf(`SELECT * FROM users;`)
	rows, err := db.Query(query)
	handleError(err)
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&iddb, &u.FName, &u.LName, &u.ID, &u.Password)
		handleError(err)
		users = append(users, u)
	}
	return users
}

func tryLogIn(un string, password string) bool {

	var p string
	query := fmt.Sprintf(`SELECT password FROM users WHERE userid="%s";`, un)
	err = db.QueryRow(query).Scan(&p)

	if err != nil {
		fmt.Println(err)
		return false
	}

	if bcrypt.CompareHashAndPassword([]byte(p), []byte(password)) != nil {
		return false
	}

	return true
}

//Session's func

func newSession(un string, cookieValue string) {

	query := fmt.Sprintf(`INSERT INTO dbsession (userid, uuid) VALUES ("%s", "%s");`, un, cookieValue)
	stmt, err := db.Prepare(query)
	handleError(err)
	defer stmt.Close()

	_, err = stmt.Exec()
	handleError(err)
}

func deleteSession(cookieValue string) {
	query := fmt.Sprintf(`DELETE FROM dbsession WHERE uuid="%s";`, cookieValue)
	stmt, err := db.Prepare(query)
	handleError(err)
	defer stmt.Close()

	_, err = stmt.Exec()
	handleError(err)
}

func getIDFromValue(cookieValue string) string {

	var userid string

	query := fmt.Sprintf(`SELECT userid FROM dbsession WHERE uuid="%s";`, cookieValue)
	err = db.QueryRow(query).Scan(&userid)
	handleError(err)

	return userid
}
