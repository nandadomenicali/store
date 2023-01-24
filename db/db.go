package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectionDb() *sql.DB {
	conection := "user=postgres dbname=store password=database sslmode=disable"
	db, err := sql.Open("postgres", conection)

	if err != nil {
		panic(err.Error())
	}
	return db
}
