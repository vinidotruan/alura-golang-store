package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func DatabaseCon() *sql.DB {
	connection := "user=root dbname=storego password=root host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)

	if err != nil {
		panic(err.Error())
	}

	return db
}
