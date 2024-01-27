-- +goose Up
ALTER TABLE grammars
ADD COLUMN difficulty VARCHAR(64);

-- +goose Down
ALTER TABLE grammars DROP COLUMN difficulty;