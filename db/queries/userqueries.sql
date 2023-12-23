-- name: InsertUser :one
Insert into users (email, hashed_password) values ($1, $2) returning id;

