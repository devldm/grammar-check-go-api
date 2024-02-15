-- +goose Up
ALTER TABLE users
ALTER COLUMN clerk_email SET NOT NULL,
ALTER COLUMN clerk_username SET NOT NULL;

-- +goose Down  
ALTER TABLE users
ALTER COLUMN clerk_email DROP NOT NULL,
ALTER COLUMN clerk_username DROP NOT NULL;
