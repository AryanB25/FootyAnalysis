CREATE TABLE IF NOT EXISTS players (
    player_id INTEGER PRIMARY KEY,

    short_name TEXT,
    long_name TEXT,
    age INTEGER,
    dob DATE,
    nationality_name TEXT,

    club_name TEXT,
    league_name TEXT,
    league_level INTEGER,
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

    attacking_finishing INTEGER,
    attacking_short_passing INTEGER,
    skill_dribbling INTEGER,
    skill_ball_control INTEGER,
    power_shot_power INTEGER,
    power_stamina INTEGER,
    power_strength INTEGER,

    mentality_aggression INTEGER,
    mentality_positioning INTEGER,
    mentality_vision INTEGER,
    mentality_penalties INTEGER,
    mentality_composure INTEGER,

    goalkeeping_diving INTEGER,
    goalkeeping_handling INTEGER,
    goalkeeping_kicking INTEGER,
    goalkeeping_positioning INTEGER,
    goalkeeping_reflexes INTEGER
);