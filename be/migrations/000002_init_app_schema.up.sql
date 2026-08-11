CREATE TABLE IF NOT EXISTS movie_series (
    ms_id SERIAL PRIMARY KEY,
    title VARCHAR(100),
    synopsis TEXT NOT NULL,
    release_year INT NOT NULL,
    classification_ms VARCHAR(3) NOT NULL,
    director VARCHAR(100),
    cover TEXT NOT NULL,
    date_added DATE,
    duration REAL,
    ms_type VARCHAR(5),
    trailer VARCHAR(300)
);

CREATE TABLE IF NOT EXISTS genres (
    genre_id SERIAL PRIMARY KEY,
    genre_name VARCHAR(100)
);

CREATE TABLE IF NOT EXISTS actors (
    actor_id SERIAL PRIMARY KEY,
    actor_first_name VARCHAR(100),
    actor_last_name VARCHAR(100),
    actor_photo TEXT NULL
);

CREATE TABLE IF NOT EXISTS movie_serie_genres (
    ms_id INT NOT NULL REFERENCES movie_series(ms_id),
    genre_id INT NOT NULL REFERENCES genres(genre_id),
    PRIMARY KEY (ms_id, genre_id)
);

CREATE TABLE IF NOT EXISTS movie_serie_actors (
    ms_id INT NOT NULL REFERENCES movie_series(ms_id),
    actor_id INT NOT NULL REFERENCES actors(actor_id),
    PRIMARY KEY (ms_id, actor_id)
);

CREATE TABLE IF NOT EXISTS episodes (
    ms_id INT NOT NULL REFERENCES movie_series(ms_id),
    season_id INT NOT NULL,
    episode_id INT NOT NULL,
    title VARCHAR(100),
    duration REAL NOT NULL,
    PRIMARY KEY (ms_id, season_id, episode_id)
);

CREATE TABLE IF NOT EXISTS comments (
    comment_id SERIAL PRIMARY KEY,
    comment TEXT,
    app_user VARCHAR(256),
    ms_id INT NOT NULL REFERENCES movie_series(ms_id)
);

CREATE TABLE IF NOT EXISTS ratings (
    rating_id SERIAL PRIMARY KEY,
    rating INT,
    app_user VARCHAR(256),
    ms_id INT NOT NULL REFERENCES movie_series(ms_id)
);

CREATE OR REPLACE FUNCTION ObtenerUltimas10Peliculas()
RETURNS SETOF movie_series
LANGUAGE plpgsql
AS $BODY$
BEGIN
    RETURN QUERY
    SELECT ms_id, title, synopsis, release_year, classification_ms,
           director, cover, date_added, duration, ms_type, trailer
    FROM movie_series
    ORDER BY ms_id DESC
    LIMIT 10;
END;
$BODY$;
