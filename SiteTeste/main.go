package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var tpl *template.Template
var db *sql.DB
var err error

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type user struct {
	FName    string
	LName    string
	ID       string
	Password string
}

func main() {

	db, err = sql.Open("mysql", "root:123456@tcp(localhost:3306)/gotest?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}

	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("css"))))

	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)

	http.ListenAndServe(":8080", nil)

}

func alreadyLoggedIn(req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}

	return getUserFromCookie(c.Value)

}

func handleError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
