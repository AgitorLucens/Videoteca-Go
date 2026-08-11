CREATE TABLE ratings (
    rating_id SERIAL PRIMARY KEY,
    rating INT,
    app_user VARCHAR(256),
    ms_id INT NOT NULL REFERENCES movie_series(ms_id)
);
