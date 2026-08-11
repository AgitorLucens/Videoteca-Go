CREATE TABLE roles (
       role_id INT PRIMARY KEY,
       role_name VARCHAR(10) UNIQUE NOT NULL
   );


INSERT INTO roles (role_id, role_name) VALUES
    (1, 'superadmin'),
    (2, 'admin'),
    (3, 'user');