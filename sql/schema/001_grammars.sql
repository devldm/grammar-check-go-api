-- +goose Up
CREATE TABLE grammars (
    id uuid PRIMARY KEY NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    grammar TEXT UNIQUE NOT NULL,
    description TEXT
);

-- +goose Down
DROP TABLE grammars;