-- +goose Up
CREATE TABLE posts (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    published_at TIMESTAMP,
    title VARCHAR(40) NOT NULL,
    url  VARCHAR(80) UNIQUE NOT NULL,
    description  VARCHAR(280),
    feed_id UUID NOT NULL, 
    FOREIGN KEY(feed_id) REFERENCES feeds(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE posts;