package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://tsvaeman:Ph2pmfh4KfKF2ReSCm9jllWPJ-thhI04@john.db.elephantsql.com:5432/tsvaeman")
	// db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
}

func Conn() *sql.DB {
	return db
}
