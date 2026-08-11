CREATE TABLE user_roles (
    user_id INT,
    role_id INT,
    PRIMARY KEY (user_id, role_id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (role_id) REFERENCES roles(role_id) ON DELETE CASCADE
);

-- Assign 'superadmin' role to user with id 1
INSERT INTO user_roles (user_id, role_id) VALUES (1, 1);
-- Assign 'admin' role to user with id 2
INSERT INTO user_roles (user_id, role_id) VALUES (1, 2); 
-- Assign 'user' role to user with id 2
INSERT INTO user_roles (user_id, role_id) VALUES (2, 3);