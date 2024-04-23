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

-- name: UpdateProfileBank :exec
UPDATE profile SET bank = $2
WHERE id = $1;

-- name: UpdateProfileBankAccountNumber :exec
UPDATE profile SET bank_account_number = $2
WHERE id = $1;

-- name: UpdateProfileClearingNumber :exec
UPDATE profile SET clearing_number = $2
WHERE id = $1;
