CREATE TABLE movie_serie_genres (
    ms_id INT NOT NULL REFERENCES movie_series(ms_id),
    genre_id INT NOT NULL REFERENCES genres(genre_id),
    PRIMARY KEY (ms_id, genre_id)
);
