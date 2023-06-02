CREATE TABLE IF NOT EXISTS users (
    id serial PRIMARY KEY,
    name varchar NOT NULL ,
    email varchar NOT NULL ,
    password varchar NOT NULL ,
    role_id int NOT NULL,
    created_at timestamp NULL DEFAULT NOW(),
    updated_at timestamp NULL DEFAULT NOW(),
    CONSTRAINT email_unique UNIQUE (email),
    CONSTRAINT fk_roles
        FOREIGN KEY(role_id)
            REFERENCES roles(id)
);

-- Bcrypt Hash is password
INSERT INTO "users" ("name", "email", "password", "role_id") VALUES ('Admin', 'admin@gmail.com', '$2a$10$KtPu3xffLGo1QAsLHw05/.dDzBVKCH3DhIT5aw9o32/rv7tMHUzDm', 1);