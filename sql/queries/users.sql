-- name: GetUsers :many
SELECT * FROM users
LIMIT $1;

-- name: GetUserById :one
SELECT * FROM users 
WHERE id = $1;

-- name: GetUserByClerkId :one
SELECT * FROM users 
WHERE clerk_id = $1;

-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, clerk_id)
VALUES ($1, $2, $3, $4)
RETURNING *;