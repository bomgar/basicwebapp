-- name: InsertUser :one
Insert into users (email, hashed_password) values ($1, $2) returning id;

-- name: SelectPasswordHashByUserEmail :one
SELECT id, hashed_password FROM users WHERE email = $1;

