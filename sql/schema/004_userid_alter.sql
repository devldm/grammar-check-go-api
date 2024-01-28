-- +goose Up
ALTER TABLE users
ADD COLUMN clerk_id VARCHAR(64);

-- +goose Down
ALTER TABLE users DROP COLUMN clerk_id;