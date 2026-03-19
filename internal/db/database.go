package db

import (
	"database/sql"
	"log"
)

func ConnectToDb() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@db:5432/barber_shop?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}
