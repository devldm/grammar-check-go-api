-- +goose Up
CREATE TABLE users (
    id uuid PRIMARY KEY NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE users;