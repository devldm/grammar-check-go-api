-- name: GetSolutions :many
SELECT * FROM solutions
LIMIT $1;

-- name: GetSolutionsByUserId :many
SELECT * FROM solutions 
WHERE user_id = $1;

-- name: GetHasUserSolved :one
SELECT * FROM solutions
WHERE user_id = $1 AND grammar_id = $2
LIMIT 1;

-- name: CreateSolution :one
INSERT INTO solutions (id, created_at, updated_at, grammar_id, user_id, solution, grammar)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;