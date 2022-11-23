package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectDatabase() *sql.DB {
	conect := "user=postgres dbname=alura_store password=111721 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conect)
	if err != nil {
		panic(err.Error())
	}
	return db
}
