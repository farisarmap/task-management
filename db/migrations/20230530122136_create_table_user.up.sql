CREATE TABLE IF NOT EXISTS users (
    id serial PRIMARY KEY,
    name VARCHAR NOT NULL,
    password VARCHAR NOT NULL,
    email VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
)