-- +goose Up
CREATE TABLE solutions (
    id uuid PRIMARY KEY NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    grammar_id UUID NOT NULL REFERENCES grammars(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    solution TEXT NOT NULL
);

-- +goose Down
DROP TABLE solutions;
