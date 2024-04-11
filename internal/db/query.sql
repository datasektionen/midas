-- name: GetProfile :one
SELECT * FROM profile
WHERE kth_id = $1 LIMIT 1;

-- name: ListProfiles :many
SELECT * FROM profile
ORDER BY kth_id;

-- name: CreateProfile :one
iNSERT INTO profile (
    kth_id
) VALUES (
    $1
)
RETURNING *;
