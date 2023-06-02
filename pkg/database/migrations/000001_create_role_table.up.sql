CREATE TABLE IF NOT EXISTS roles (
    id serial PRIMARY KEY,
    name varchar NOT NULL ,
    created_at timestamp NULL DEFAULT NOW(),
    updated_at timestamp NULL DEFAULT NOW()
);

INSERT INTO "roles" ("name") VALUES ('admin');
INSERT INTO "roles" ("name") VALUES ('user');