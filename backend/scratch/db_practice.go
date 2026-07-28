package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {
	currStr := "host=localhost port=5432 user=postgres password=pass dbname=football sslmode=disable"
	database, err := sql.Open("postgres", currStr)
	if err != nil {
		fmt.Println(err)
	}
	_, err = database.Exec("CREATE TABLE IF NOT EXISTS scratch_test (id SERIAL PRIMARY KEY, note TEXT)")
	if err != nil {
		fmt.Println(err)
	}
	_, err = database.Exec("INSERT INTO scratch_test (note) VALUES ($1)", "hello") 
	if err != nil {
		fmt.Println(err)
	}
	rows, err := database.Query("SELECT note FROM scratch_test")
	if err != nil {
		fmt.Println(err)
	}
	var number string
	var currentString []string
	for rows.Next() {
		rows.Scan(&number)
		currentString = append(currentString, number)
	}
	fmt.Println(currentString)
}
