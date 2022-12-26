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

	stmt, err := db.Prepare("UPDATE users SET name=$2 WHERE id=$1;")

	if err != nil {
		log.Fatal("can't prepare statement update", err)
	}

	if _, err := stmt.Exec(1, "update-user"); err != nil {
		log.Fatal("error execute update ", err)
	}

	fmt.Println("update success")
}
