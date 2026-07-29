package seed

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func toIntOrZero(s string) int {
	if s == "" {
		return 0
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return n
}

func toInt64OrZero(s string) int64 {
	if s == "" {
		return 0
	}
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return n
}

func toFloatOrZero(s string) float64 {
	if s == "" {
		return 0
	}

	n, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}

	return n
}

func Run(database *sql.DB) error {
	_, err := database.Exec(`
		CREATE TABLE IF NOT EXISTS players (
			player_id INTEGER PRIMARY KEY,
			short_name TEXT,
			long_name TEXT,
			age INTEGER,
			nationality_name TEXT,
			club_name TEXT,
			league_name TEXT,
			club_position TEXT,
			height_cm INTEGER,
			weight_kg INTEGER,
			preferred_foot TEXT,
			value_eur BIGINT,
			wage_eur BIGINT,
			overall INTEGER,
			potential INTEGER,
			international_reputation INTEGER,
			pace INTEGER,
			shooting INTEGER,
			passing INTEGER,
			dribbling INTEGER,
			defending INTEGER,
			physic INTEGER
		)
	`)
	if err != nil {
		return err
	}

	var count int
	err = database.QueryRow("SELECT COUNT(*) FROM players").Scan(&count)
	if err == nil && count > 0 {
		return nil
	}

	file, err := os.Open("data/male_players.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	header, err := reader.Read()
	if err != nil {
		return err
	}

	colIndex := make(map[string]int)
	for i, colName := range header {
		colIndex[colName] = i
	}

	rowsAll, err := reader.ReadAll()
	if err != nil {
		return err
	}

	latestVersion := 0
	for _, row := range rowsAll {
		v := int(toFloatOrZero(row[colIndex["fifa_version"]]))
		if v > latestVersion {
			latestVersion = v
		}
	}

	fmt.Println("Latest FIFA version:", latestVersion)

	latestCount := 0
	for _, row := range rowsAll {
		if int(toFloatOrZero(row[colIndex["fifa_version"]])) == latestVersion {
			latestCount++
		}
	}

	fmt.Println("Rows with latest FIFA version:", latestCount)

	columns := []string{
		"player_id", "short_name", "long_name", "age", "nationality_name",
		"club_name", "league_name", "club_position", "height_cm", "weight_kg",
		"preferred_foot", "value_eur", "wage_eur", "overall", "potential",
		"international_reputation", "pace", "shooting", "passing",
		"dribbling", "defending", "physic",
	}

	placeholders := make([]string, len(columns))
	for i := range columns {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
	}

	insertSQL := fmt.Sprintf(
		"INSERT INTO players (%s) VALUES (%s) ON CONFLICT (player_id) DO NOTHING",
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "),
	)

	inserted := 0
	for _, row := range rowsAll {
		if int(toFloatOrZero(row[colIndex["fifa_version"]])) != latestVersion {
			continue
		}

		_, err := database.Exec(insertSQL,
			toIntOrZero(row[colIndex["player_id"]]),
			row[colIndex["short_name"]],
			row[colIndex["long_name"]],
			toIntOrZero(row[colIndex["age"]]),
			row[colIndex["nationality_name"]],
			row[colIndex["club_name"]],
			row[colIndex["league_name"]],
			row[colIndex["club_position"]],
			toIntOrZero(row[colIndex["height_cm"]]),
			toIntOrZero(row[colIndex["weight_kg"]]),
			row[colIndex["preferred_foot"]],
			int64(toFloatOrZero(row[colIndex["value_eur"]])),
			int64(toFloatOrZero(row[colIndex["wage_eur"]])),
			toIntOrZero(row[colIndex["overall"]]),
			toIntOrZero(row[colIndex["potential"]]),
			toIntOrZero(row[colIndex["international_reputation"]]),
			toIntOrZero(row[colIndex["pace"]]),
			toIntOrZero(row[colIndex["shooting"]]),
			toIntOrZero(row[colIndex["passing"]]),
			toIntOrZero(row[colIndex["dribbling"]]),
			toIntOrZero(row[colIndex["defending"]]),
			toIntOrZero(row[colIndex["physic"]]),
		)
		if err != nil {
			return err
		}
		inserted++
	}

	fmt.Println("Seed complete!")
	fmt.Println("Inserted:", inserted)
	fmt.Println("FIFA version:", latestVersion)
	return nil
}
