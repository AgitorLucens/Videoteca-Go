CREATE TABLE episodes (
    ms_id INT NOT NULL REFERENCES movie_series(ms_id),
    season_id INT NOT NULL,
    episode_id INT NOT NULL,
    title VARCHAR(100),
    duration REAL NOT NULL,
    PRIMARY KEY (ms_id, season_id, episode_id)
);