-- name: CreateUserWithProfile :one
WITH new_user AS (
    INSERT INTO users (id, username, email, password_hash)
    VALUES ($1, $2, $3, $4)
    RETURNING id
)
INSERT INTO user_profiles (user_id, first_name, last_name, phone_number, address)
SELECT id, $5, $6, $7, $8
FROM new_user
RETURNING user_id, first_name, last_name, phone_number, address;

-- name: GetUserByUsernameOrEmail :one
SELECT id, username, email, password_hash, created_at, updated_at
FROM users
WHERE username = $1 OR email = $1;

-- name: GetUserByID :one
SELECT u.id, u.username, u.email, u.created_at, u.updated_at,
       p.first_name, p.last_name, p.phone_number, p.address
FROM users u
JOIN user_profiles p ON u.id = p.user_id
WHERE u.id = $1;

-- name: UpdateUserProfile :exec
UPDATE user_profiles
SET first_name = $2, last_name = $3, phone_number = $4, address = $5, updated_at = now()
WHERE user_id = $1;

-- name: DeleteUserAndProfile :exec
DELETE FROM user_profiles
WHERE user_id = $1;

DELETE FROM users
WHERE id = $1;