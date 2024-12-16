-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES (
        gen_random_uuid(),
        NOW(),
        NOW(),
        $1,
        $2
)
RETURNING *)
SELECT
    inserted_feed_follow.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted_feed_follow
JOIN feeds ON inserted_feed_follow.feed_id = feeds.id
JOIN users ON inserted_feed_follow.user_id = users.id;




-- name: GetFeedFollowsForUser :many
SELECT *, feeds.name AS feed_name 
FROM feed_follows
JOIN feeds ON feed_id = feeds.id
JOIN users ON user_id = users.id
WHERE users.name = $1;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows;