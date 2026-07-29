package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("male_players.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	header, err := reader.Read()
	if err != nil {
		fmt.Println(err)
		return
	}
	
	rows := make(map[string]int)

	for i, colName := range header {
		rows[colName] = i
	}

	for i := range rows {
		fmt.Println(rows[i])
	}

		// confirm the map has the values we expect
	fmt.Println("short_name index:", rows["short_name"])
	fmt.Println("overall index:", rows["overall"])
	fmt.Println("pace index:", rows["pace"])

	// read the next 3 data rows and print two fields from each
	for i := 0; i < 3; i++ {
		row, err := reader.Read()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(row[rows["short_name"]], "-", row[rows["overall"]])
	}

}