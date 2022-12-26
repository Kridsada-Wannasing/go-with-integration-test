package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	url := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", url)

	if err != nil {
		log.Fatal("Connect to database error", err)
	}

	defer db.Close()

	stmt, err := db.Prepare("SELECT id, name, age FROM users where id=$1")
	if err != nil {
		log.Fatal("can't prepare query one user statement", err)
	}

	rowId := 1
	row := stmt.QueryRow(rowId)

	var id int
	var name string
	var age int
	err = row.Scan(&id, &name, &age)
	if err != nil {
		log.Fatal("can't Scan row into variable", err)
	}
	fmt.Println("onw row", id, name, age)
}
