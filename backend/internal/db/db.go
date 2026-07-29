package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	connectionString := "host=localhost port=5432 user=postgres password=pass dbname=football sslmode=disable"
	database, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	err = database.Ping()
	if err != nil {
		return nil, err
	}
	return database, err
}
