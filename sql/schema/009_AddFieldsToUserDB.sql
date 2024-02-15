-- +goose Up
ALTER TABLE users
ADD COLUMN clerk_email VARCHAR(64),
ADD COLUMN clerk_username VARCHAR(64);


-- +goose Down
ALTER TABLE users 
DROP COLUMN clerk_email,
DROP COLUMN clerk_username;