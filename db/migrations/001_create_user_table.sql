-- +goose Up
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email text NOT NULL UNIQUE,
    hashed_password text NOT NULL
);
