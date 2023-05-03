CREATE TABLE users (
    "id"            SERIAL PRIMARY KEY,
    "email"         VARCHAR(255) NOT NULL UNIQUE,
    "password_hash" VARCHAR(255) NOT NULL,
    "first_name"    VARCHAR(25) NOT NULL,
    "last_name"     VARCHAR(55) NOT NULL
);
