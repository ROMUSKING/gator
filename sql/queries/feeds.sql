-- name: CreateFeed :one

INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1,
    $2,
    $3 
)
RETURNING *;

-- name: GetFeed :one
SELECT * FROM feeds
WHERE url = $1;

-- name: GetFeeds :many
SELECT 
    feeds.name as feed_name, 
    feeds.url as url, 
    users.name as user_name 
FROM feeds
JOIN users ON users.id = feeds.user_id;

-- name: DeleteFeeds :exec
DELETE FROM feeds;