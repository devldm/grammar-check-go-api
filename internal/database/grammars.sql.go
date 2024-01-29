// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: grammars.sql

package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createGrammar = `-- name: CreateGrammar :one
INSERT INTO grammars (id, created_at, updated_at, grammar, description, difficulty)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, created_at, updated_at, grammar, description, difficulty
`

type CreateGrammarParams struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Grammar     string
	Description sql.NullString
	Difficulty  sql.NullString
}

func (q *Queries) CreateGrammar(ctx context.Context, arg CreateGrammarParams) (Grammar, error) {
	row := q.db.QueryRowContext(ctx, createGrammar,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Grammar,
		arg.Description,
		arg.Difficulty,
	)
	var i Grammar
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Grammar,
		&i.Description,
		&i.Difficulty,
	)
	return i, err
}

const getGrammarById = `-- name: GetGrammarById :one
SELECT id, created_at, updated_at, grammar, description, difficulty FROM grammars 
WHERE id = $1
`

func (q *Queries) GetGrammarById(ctx context.Context, id uuid.UUID) (Grammar, error) {
	row := q.db.QueryRowContext(ctx, getGrammarById, id)
	var i Grammar
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Grammar,
		&i.Description,
		&i.Difficulty,
	)
	return i, err
}

const getGrammars = `-- name: GetGrammars :many
SELECT id, created_at, updated_at, grammar, description, difficulty FROM grammars
LIMIT $1
`

func (q *Queries) GetGrammars(ctx context.Context, limit int32) ([]Grammar, error) {
	rows, err := q.db.QueryContext(ctx, getGrammars, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Grammar
	for rows.Next() {
		var i Grammar
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Grammar,
			&i.Description,
			&i.Difficulty,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}