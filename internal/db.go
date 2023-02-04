package internal

import (
	"database/sql"
	"log"
)

func NewDb() {

}

// подключение к БД
func Connection() *sql.DB {
	conString := "host=localhost port=5432 user=postgres password=10unibos dbname=users sslmode=disable"

	conn, err := sql.Open("postgres", conString)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}
