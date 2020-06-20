package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	createTableIfNotExist()
}

func createTableIfNotExist() {
	createTable := `CREATE TABLE IF NOT EXISTS customers (
		id SERIAL PRIMARY KEY,
		name TEXT,
		email TEXT,
		status TEXT
	);`

	if _, err := db.Exec(createTable); err != nil {
		log.Fatal("Can't create table", err)
	}
}

//Conn
func Conn() *sql.DB {
	return db
}
