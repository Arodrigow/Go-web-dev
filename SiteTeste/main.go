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
	Password []byte
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

func index(w http.ResponseWriter, req *http.Request) {
	err = tpl.ExecuteTemplate(w, "index.gohtml", nil)
	handleError(w, err)
}

func login(w http.ResponseWriter, req *http.Request) {
	err = tpl.ExecuteTemplate(w, "login.gohtml", nil)
	handleError(w, err)
}

func handleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
