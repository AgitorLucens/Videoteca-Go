CREATE TABLE comments (
    comment_id SERIAL PRIMARY KEY,
    comment TEXT,
    app_user VARCHAR(256),
    ms_id INT NOT NULL REFERENCES movie_series(ms_id)
);