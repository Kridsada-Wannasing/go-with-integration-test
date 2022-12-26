package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/Kridsada-Wannasing/gosql/kridsada" // If it has been imported, the func init() will be called

	_ "github.com/lib/pq" // required for init sql driver
)

func init() {
	fmt.Println("main init")
} // this func seems like constructor

func main() {
	kridsada.Say()

	url := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", url)

	if err != nil {
		log.Fatal("Connect to database error", err)
	}

	defer db.Close()

	createTb := `
	CREATE TABLE IF NOT EXISTS users ( id SERIAL PRIMARY KEY, name TEXT, age INT );
	`

	_, err = db.Exec(createTb)

	if err != nil {
		log.Fatal("can't create table", err)
	}

	fmt.Println("create table success")

	log.Println("ok")
}
