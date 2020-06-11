package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	FName    string
	LName    string
	ID       string
	Password string
}

func main() {

	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/gotest?charset=utf8")
	testeError(err)
	defer db.Close()

	err = db.Ping()
	testeError(err)

	//SELECTING A SINGLE ROW
	var usuario user
	query := fmt.Sprintf(`SELECT * FROM users WHERE id="Euz";`)
	err = db.QueryRow(query).Scan(&usuario.FName, &usuario.LName, &usuario.ID, &usuario.Password)
	//RETURNS A ERROR IF ID IS NON-EXISTENT, AND MAIN CONTINUES TO RUN
	testeError(err)
	fmt.Println("QueryRow: ", usuario)

	//SELECTING ALL ROWS FROM THE TABLE USERS
	var usuarios []user
	query = fmt.Sprintf(`SELECT * FROM users;`)
	rows, err := db.Query(query)
	testeError(err)
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&usuario.FName, &usuario.LName, &usuario.ID, &usuario.Password)
		testeError(err)
		usuarios = append(usuarios, usuario)
	}
	fmt.Println()
	fmt.Println("Query: ", usuarios)

	//SELECTING A COLUMN FROM THE TABLE USERS

	var column string
	var columns []string

	query = fmt.Sprintf(`SELECT id FROM users;`)
	rows, err = db.Query(query)
	testeError(err)
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&column)
		testeError(err)
		columns = append(columns, column)
	}
	fmt.Println()
	fmt.Println("Column query: ", columns)

	//INSERTING VALUES TO TABLE USERS
	stmt, err := db.Prepare(`INSERT INTO users (fName, lName, id, password) VALUES("Rodrigo", "Alves", "R", "123456");`)
	testeError(err)
	defer stmt.Close()

	_, err = stmt.Exec()
	testeError(err)

	//DELETING VALUES TO TABLE USERS
	stmt, err = db.Prepare(`UPDATE users SET id="Gopher" WHERE id="R"`)
	testeError(err)
	defer stmt.Close()

	_, err = stmt.Exec()
	testeError(err)

	//DELETING VALUES TO TABLE USERS
	stmt, err = db.Prepare(`DELETE FROM users WHERE id="Gopher";`)
	testeError(err)
	defer stmt.Close()

	_, err = stmt.Exec()
	testeError(err)
}

func testeError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
