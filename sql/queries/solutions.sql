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

-- name: GetSolutionsWithUserData :many
SELECT s.*, u.clerk_username, u.clerk_email, u.clerk_image, u.id AS user_id
FROM solutions s
JOIN users u ON s.user_id = u.id
LIMIT $1;

-- name: GetSolutionsByGrammarIdWithUserData :many
SELECT s.*, u.clerk_username, u.clerk_email, u.clerk_image, u.id AS user_id
FROM solutions s
JOIN users u ON s.user_id = u.id
WHERE s.grammar_id = $1
LIMIT $2;

-- name: DeleteSolutionBySolutionId :exec
DELETE FROM solutions
WHERE id = $1;
