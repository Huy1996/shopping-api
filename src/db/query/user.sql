-- name: CreateUserCredential :one
INSERT INTO user_credential (
    id,
    username,
    hashed_password,
    email
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetUserCredential :one
SELECT * FROM user_credential
WHERE username = $1
LIMIT 1;

-- name: CreateUserInfo :one
INSERT INTO user_info (
    id,
    user_id,
    phone_number,
    first_name,
    last_name,
    middle_name
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING  *;

-- name: GetUserInfoByID :one
SELECT * FROM user_info
WHERE id = $1
LIMIT 1;

-- name: GetUserInfoByUserID :one
SELECT * FROM user_info
WHERE user_id = $1
LIMIT 1;

-- name: CreateAddressBook :one
INSERT INTO address_book (
    id,
    owner,
    address_name,
    address,
    city,
    state,
    zipcode
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: GetListAddresses :many
SELECT * FROM address_book
WHERE owner = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: GetNumberAddresses :one
SELECT COUNT(*) FROM address_book
WHERE owner = $1
LIMIT 1;