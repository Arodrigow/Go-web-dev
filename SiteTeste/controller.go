package main

import (
	"fmt"
	"log"
)

//USER CONTROLLER
func userExist(value string) bool {
	_, err := db.Query(`SELECT id FROM users WHERE id=` + value + `;`)
	if err != nil {
		return false
	}

	return true
}

func addUser(ID string, fName string, lName string, password string) {
	querie := fmt.Sprintf(`INSERT INTO users (fName, lName, id, password) VALUES(" %s", "%s", "%s", "%s");`, fName, lName, ID, password)
	stmt, err := db.Prepare(querie)
	if err != nil {
		log.Println(err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		log.Println(err)
		return
	}
}

//SESSION CONTROLER
func addSession(ID string, UUID string) {
	querie := fmt.Sprintf(`INSERT INTO dbsession (userid, uuid) VALUES("%s", "%s");`, ID, UUID)
	stmt, err := db.Prepare(querie)
	if err != nil {
		log.Println("EH AQUI O ERRO")
		log.Println(err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		log.Println(err)
		return
	}
}
