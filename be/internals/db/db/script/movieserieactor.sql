CREATE TABLE movie_serie_actors (
    ms_id INT NOT NULL REFERENCES movie_series(ms_id),
    actor_id INT NOT NULL REFERENCES actors(actor_id),
    PRIMARY KEY (ms_id, actor_id)
);