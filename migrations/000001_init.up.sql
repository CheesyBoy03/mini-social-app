CREATE TABLE users (
    id            SERIAL PRIMARY KEY,
    email         VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    first_name    VARCHAR(25) NOT NULL,
    last_name     VARCHAR(55) NOT NULL
);

CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    author_id INT NOTE NULL,
    title VARCHAR(70) NOT NULL,
    description VARCHAR(255),
    published_at timestamp DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (author_id),
    FOREIGN KEY (author_id) REFERENCES users(id) ON DELETE CASCADE,
)
