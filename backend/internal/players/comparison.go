package players

type PlayerComparison struct {
	Player1 PlayerComparisonStats `json:"player1"`
	Player2 PlayerComparisonStats `json:"player2"`
}

type PlayerComparisonStats struct {
	Name        string `json:"name"`
	Age         int    `json:"age"`
	Nationality string `json:"nationality"`

	Club     string `json:"club"`
	League   string `json:"league"`
	Position string `json:"position"`

	Overall   int `json:"overall"`
	Potential int `json:"potential"`

	Pace      int `json:"pace"`
	Shooting  int `json:"shooting"`
	Passing   int `json:"passing"`
	Dribbling int `json:"dribbling"`
	Defending int `json:"defending"`
	Physical  int `json:"physical"`
}

func convertToComparisonStats(player Player) PlayerComparisonStats {
	var playerComparison PlayerComparisonStats

	playerComparison.Name = player.FullName
	playerComparison.Age = player.Age
	playerComparison.Nationality = player.Nationality

	playerComparison.Club = player.Club
	playerComparison.League = player.League
	playerComparison.Position = player.Position
	playerComparison.Overall = player.Overall
	playerComparison.Potential = player.Potential

	playerComparison.Pace = player.Pace
	playerComparison.Shooting = player.Shooting
	playerComparison.Passing = player.Passing
	playerComparison.Dribbling = player.Dribbling
	playerComparison.Defending = player.Defending
	playerComparison.Physical = player.Physical

	return playerComparison
}
