package players

type Player struct {
	PlayerID    int    `json:"player_id"`
	ShortName   string `json:"short_name"`
	FullName    string `json:"long_name"`
	Age         int    `json:"age"`
	Nationality string `json:"nationality"`

	Club     string `json:"club"`
	League   string `json:"league"`
	Position string `json:"position"`

	Overall    int `json:"overall"`
	Potential  int `json:"potential"`
	Reputation int `json:"international_reputation"`

	Pace      int `json:"pace"`
	Shooting  int `json:"shooting"`
	Passing   int `json:"passing"`
	Dribbling int `json:"dribbling"`
	Defending int `json:"defending"`
	Physical  int `json:"physical"`

	PreferredFoot string `json:"preferred_foot"`

	Height int   `json:"height_cm"`
	Weight int   `json:"weight_kg"`
	Value  int64 `json:"value_eur"`
	Wage   int64 `json:"wage_eur"`
}
