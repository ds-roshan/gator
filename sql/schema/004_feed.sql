-- +goose Up
ALTER TABLE feeds
ADD Column IF NOT EXISTS last_fetched_at TIMESTAMP;

-- +goose Down
ALTER TABLE feeds
DROP COLUMN IF EXISTS last_fetched_at;
