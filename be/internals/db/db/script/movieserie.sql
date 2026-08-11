CREATE TABLE movie_series (
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