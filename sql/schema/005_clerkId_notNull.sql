-- +goose Up
ALTER TABLE users
ALTER COLUMN clerk_id SET NOT NULL;

-- +goose Down  
ALTER TABLE users
ALTER COLUMN clerk_id DROP NOT NULL;