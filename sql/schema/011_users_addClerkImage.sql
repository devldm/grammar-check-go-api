-- +goose Up
ALTER TABLE users
ADD COLUMN clerk_image TEXT NOT NULL;

-- +goose Down
ALTER TABLE users DROP COLUMN clerk_image;