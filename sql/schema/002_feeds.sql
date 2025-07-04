-- +goose Up
CREATE TABLE feeds (
id UUID PRIMARY KEY,
name TEXT NOT NULL,
url TEXT NOT NULL UNIQUE,
created_at TIMESTAMP NOT NULL,
updated_at TIMESTAMP NOT NULL,
user_id UUID NOT NULL,
CONSTRAINT fk_users
    FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS feeds;
