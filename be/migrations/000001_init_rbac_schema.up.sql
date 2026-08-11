CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255),
    password VARCHAR(255),
    profile_picture BYTEA,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS roles (
    role_id INT PRIMARY KEY,
    role_name VARCHAR(10) UNIQUE NOT NULL
);

INSERT INTO roles (role_id, role_name) VALUES
    (1, 'superadmin'),
    (2, 'admin'),
    (3, 'user');

CREATE TABLE IF NOT EXISTS user_roles (
    user_id INT,
    role_id INT,
    PRIMARY KEY (user_id, role_id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (role_id) REFERENCES roles(role_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS access (
    access_id INT PRIMARY KEY,
    access_name VARCHAR(10)
);

INSERT INTO access (access_id, access_name) VALUES
    (1, 'create'),
    (2, 'read'),
    (3, 'update'),
    (4, 'delete');

CREATE TABLE IF NOT EXISTS role_access (
    role_id INT,
    access_id INT,
    PRIMARY KEY (role_id, access_id),
    FOREIGN KEY (role_id) REFERENCES roles(role_id),
    FOREIGN KEY (access_id) REFERENCES access(access_id)
);

INSERT INTO role_access (role_id, access_id) VALUES
    (1, 1),
    (1, 2),
    (1, 3),
    (1, 4),
    (2, 2),
    (2, 3),
    (3, 2);
