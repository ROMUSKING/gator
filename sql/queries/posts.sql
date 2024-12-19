-- name: CreatePost :exec

INSERT INTO posts (id, created_at, updated_at, published_at, title, url, description, feed_id)
VALUES (
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1,
    $2,
    $3,
    $4,
    $5 
);

-- name: GetPost :one
SELECT * FROM posts
WHERE url = $1;

-- name: GetPosts :many
SELECT * FROM posts
WHERE feed_id IN (
    SELECT feed_id FROM feed_follows
    WHERE user_id = $1)
ORDER BY published_at desc
LIMIT $2;

-- name: DeletePosts :exec
DELETE FROM posts
WHERE id = $1;
