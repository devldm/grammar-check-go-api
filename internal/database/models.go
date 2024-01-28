// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Grammar struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Grammar     string
	Description sql.NullString
	Difficulty  sql.NullString
}

type Solution struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	GrammarID uuid.UUID
	UserID    uuid.UUID
	Solution  string
	Grammar   string
}

type User struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	ClerkID   string
}
