package user

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() {
	var err error
	url := os.Getenv("DATABASE_URL")
	db, err = sql.Open("postgres", url)

	if err != nil {
		log.Fatal("Connect to database error", err)
	}

	createTb := `
	CREATE TABLE IF NOT EXISTS users ( id SERIAL PRIMARY KEY, name TEXT, age INT );
	`

	_, err = db.Exec(createTb)

	if err != nil {
		log.Fatal("can't create table", err)
	}
}
