-- +goose Up
-- SQL in section 'Up' is executed when the migration is applied
CREATE TABLE IF NOT EXISTS posts (
    id SERIAL PRIMARY KEY,
    user_id INT,
    content TEXT,
    allow_comments BOOLEAN
);

-- +goose Down
-- SQL section 'Down' is executed when the migration is rolled back
DROP TABLE IF EXISTS posts;
