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

func (r *Repository) GetAllPlayers(page int, limit int, min_rating int, max_rating int) ([]Player, error) {
	offset := (page - 1) * limit
	var rows *sql.Rows
	var err error
	if min_rating > 0 && max_rating > 0 {
		rows, err = r.db.Query("SELECT player_id, short_name, long_name, age, nationality_name, club_name, league_name, club_position, height_cm, weight_kg, preferred_foot, value_eur, wage_eur, overall, potential, international_reputation, pace, shooting, passing, dribbling, defending, physic FROM players WHERE overall >= $1 AND overall <= $2 LIMIT $3 OFFSET $4", min_rating, max_rating, limit, offset)
	} else if min_rating > 0 && max_rating == 0 {
		rows, err = r.db.Query("SELECT player_id, short_name, long_name, age, nationality_name, club_name, league_name, club_position, height_cm, weight_kg, preferred_foot, value_eur, wage_eur, overall, potential, international_reputation, pace, shooting, passing, dribbling, defending, physic FROM players WHERE overall >= $1 LIMIT $2 OFFSET $3", min_rating, limit, offset)
	} else if min_rating == 0 && max_rating > 0 {
		rows, err = r.db.Query("SELECT player_id, short_name, long_name, age, nationality_name, club_name, league_name, club_position, height_cm, weight_kg, preferred_foot, value_eur, wage_eur, overall, potential, international_reputation, pace, shooting, passing, dribbling, defending, physic FROM players WHERE overall <= $1 LIMIT $2 OFFSET $3", max_rating, limit, offset)
	} else {
		rows, err = r.db.Query("SELECT player_id, short_name, long_name, age, nationality_name, club_name, league_name, club_position, height_cm, weight_kg, preferred_foot, value_eur, wage_eur, overall, potential, international_reputation, pace, shooting, passing, dribbling, defending, physic FROM players LIMIT $1 OFFSET $2", limit, offset)
	}
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
