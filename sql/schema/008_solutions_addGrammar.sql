-- +goose Up
ALTER TABLE solutions
ADD COLUMN grammar TEXT NOT NULL REFERENCES grammars(grammar);

-- +goose Down
ALTER TABLE solutions DROP COLUMN grammar;