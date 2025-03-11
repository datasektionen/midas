-- name: GetProfile :one
SELECT * FROM profile
WHERE ugKthid = $1 LIMIT 1;

-- name: ListProfiles :many
SELECT * FROM profile
ORDER BY ugKthid;

-- name: CreateProfile :one
iNSERT INTO profile (
    ugKthid
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
