package main

import (
	"fmt"
	"footballanalyticshub/internal/db"
)

type Player struct {
	long_name string
	age       int
	overall   int
}

func main() {
	database, err := db.Connect()
	if err != nil {
		fmt.Println(err)
		return
	}
	rows, err := database.Query("SELECT long_name, age, overall FROM players LIMIT 3")
	if err != nil {
		fmt.Println(err)
		return
	}

	var player Player
	var original_slice []Player

	for rows.Next() {
		err = rows.Scan(&player.long_name, &player.age, &player.overall)
		if err != nil {
			fmt.Println(err)
			return
		}
		original_slice = append(original_slice, player)
	}

	fmt.Print(original_slice)
}
