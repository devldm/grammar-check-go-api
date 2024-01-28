-- +goose Up
ALTER TABLE users
ADD CONSTRAINT clerk_id_unique UNIQUE (clerk_id);

-- +goose Down
ALTER TABLE users 
DROP CONSTRAINT clerk_id_unique;