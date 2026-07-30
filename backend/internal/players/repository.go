package players

import (
	"database/sql"
)

type Repository struct {
	db *sql.DB
}

func NewPlayerRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) GetAllPlayers() ([]Player, error) {
	rows, err := r.db.Query("SELECT player_id, short_name, long_name, age, nationality_name, club_name, league_name, club_position, height_cm, weight_kg, preferred_foot, value_eur, wage_eur, overall, potential, international_reputation, pace, shooting, passing, dribbling, defending, physic FROM players")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var players []Player
	for rows.Next() {
		var player Player
		err = rows.Scan(
			&player.PlayerID,
			&player.ShortName,
			&player.FullName,
			&player.Age,
			&player.Nationality,
			&player.Club,
			&player.League,
			&player.Position,
			&player.Height,
			&player.Weight,
			&player.PreferredFoot,
			&player.Value,
			&player.Wage,
			&player.Overall,
			&player.Potential,
			&player.Reputation,
			&player.Pace,
			&player.Shooting,
			&player.Passing,
			&player.Dribbling,
			&player.Defending,
			&player.Physical,
		)
		if err != nil {
			return nil, err
		}
		players = append(players, player)
	}
	return players, nil
}

func (r *Repository) GetPlayerByID(id int) (Player, error) {
	row := r.db.QueryRow("SELECT player_id, short_name, long_name, age, nationality_name, club_name, league_name, club_position, height_cm, weight_kg, preferred_foot, value_eur, wage_eur, overall, potential, international_reputation, pace, shooting, passing, dribbling, defending, physic FROM players WHERE player_id = $1", id)
	var player Player
	err := row.Scan(
		&player.PlayerID,
		&player.ShortName,
		&player.FullName,
		&player.Age,
		&player.Nationality,
		&player.Club,
		&player.League,
		&player.Position,
		&player.Height,
		&player.Weight,
		&player.PreferredFoot,
		&player.Value,
		&player.Wage,
		&player.Overall,
		&player.Potential,
		&player.Reputation,
		&player.Pace,
		&player.Shooting,
		&player.Passing,
		&player.Dribbling,
		&player.Defending,
		&player.Physical,
	)
	if err != nil {
		return player, err
	}
	return player, nil
}

func (r *Repository) SearchPlayers(name string) ([]Player, error) {
	rows, err := r.db.Query("SELECT player_id, short_name, long_name, age, nationality_name, club_name, league_name, club_position, height_cm, weight_kg, preferred_foot, value_eur, wage_eur, overall, potential, international_reputation, pace, shooting, passing, dribbling, defending, physic FROM players WHERE long_name ILIKE $1", "%"+name+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var players []Player
	for rows.Next() {
		var player Player
		err = rows.Scan(
			&player.PlayerID,
			&player.ShortName,
			&player.FullName,
			&player.Age,
			&player.Nationality,
			&player.Club,
			&player.League,
			&player.Position,
			&player.Height,
			&player.Weight,
			&player.PreferredFoot,
			&player.Value,
			&player.Wage,
			&player.Overall,
			&player.Potential,
			&player.Reputation,
			&player.Pace,
			&player.Shooting,
			&player.Passing,
			&player.Dribbling,
			&player.Defending,
			&player.Physical,
		)
		if err != nil {
			return nil, err
		}
		players = append(players, player)
	}
	return players, nil
}
