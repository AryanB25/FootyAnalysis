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
    physic INTEGER,
);