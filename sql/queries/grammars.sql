-- name: GetGrammars :many
SELECT * FROM grammars
LIMIT $1;

-- name: GetGrammarById :one
SELECT * FROM grammars 
WHERE id = $1;

-- name: CreateGrammar :one
INSERT INTO grammars (id, created_at, updated_at, grammar, description, difficulty)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;