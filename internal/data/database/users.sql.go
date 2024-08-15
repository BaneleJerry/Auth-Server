// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createUserWithProfile = `-- name: CreateUserWithProfile :one
WITH new_user AS (
    INSERT INTO users (id, username, email, password_hash)
    VALUES ($1, $2, $3, $4)
    RETURNING id
)
INSERT INTO user_profiles (user_id, first_name, last_name, phone_number, address)
SELECT id, $5, $6, $7, $8
FROM new_user
RETURNING user_id, first_name, last_name, phone_number, address
`

type CreateUserWithProfileParams struct {
	ID           uuid.UUID
	Username     string
	Email        string
	PasswordHash string
	FirstName    sql.NullString
	LastName     sql.NullString
	PhoneNumber  sql.NullString
	Address      sql.NullString
}

type CreateUserWithProfileRow struct {
	UserID      uuid.UUID
	FirstName   sql.NullString
	LastName    sql.NullString
	PhoneNumber sql.NullString
	Address     sql.NullString
}

func (q *Queries) CreateUserWithProfile(ctx context.Context, arg CreateUserWithProfileParams) (CreateUserWithProfileRow, error) {
	row := q.db.QueryRowContext(ctx, createUserWithProfile,
		arg.ID,
		arg.Username,
		arg.Email,
		arg.PasswordHash,
		arg.FirstName,
		arg.LastName,
		arg.PhoneNumber,
		arg.Address,
	)
	var i CreateUserWithProfileRow
	err := row.Scan(
		&i.UserID,
		&i.FirstName,
		&i.LastName,
		&i.PhoneNumber,
		&i.Address,
	)
	return i, err
}

const deleteUserAndProfile = `-- name: DeleteUserAndProfile :exec
DELETE FROM user_profiles
WHERE user_id = $1
`

func (q *Queries) DeleteUserAndProfile(ctx context.Context, userID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteUserAndProfile, userID)
	return err
}

const getUserByID = `-- name: GetUserByID :one
SELECT u.id, u.username, u.email, u.created_at, u.updated_at,
       p.first_name, p.last_name, p.phone_number, p.address
FROM users u
JOIN user_profiles p ON u.id = p.user_id
WHERE u.id = $1
`

type GetUserByIDRow struct {
	ID          uuid.UUID
	Username    string
	Email       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	FirstName   sql.NullString
	LastName    sql.NullString
	PhoneNumber sql.NullString
	Address     sql.NullString
}

func (q *Queries) GetUserByID(ctx context.Context, id uuid.UUID) (GetUserByIDRow, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)
	var i GetUserByIDRow
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.FirstName,
		&i.LastName,
		&i.PhoneNumber,
		&i.Address,
	)
	return i, err
}

const getUserByUsernameOrEmail = `-- name: GetUserByUsernameOrEmail :one
SELECT id, username, email, password_hash, created_at, updated_at
FROM users
WHERE username = $1 OR email = $1
`

func (q *Queries) GetUserByUsernameOrEmail(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByUsernameOrEmail, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.PasswordHash,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUserProfile = `-- name: UpdateUserProfile :exec
UPDATE user_profiles
SET first_name = $2, last_name = $3, phone_number = $4, address = $5, updated_at = now()
WHERE user_id = $1
`

type UpdateUserProfileParams struct {
	UserID      uuid.UUID
	FirstName   sql.NullString
	LastName    sql.NullString
	PhoneNumber sql.NullString
	Address     sql.NullString
}

func (q *Queries) UpdateUserProfile(ctx context.Context, arg UpdateUserProfileParams) error {
	_, err := q.db.ExecContext(ctx, updateUserProfile,
		arg.UserID,
		arg.FirstName,
		arg.LastName,
		arg.PhoneNumber,
		arg.Address,
	)
	return err
}
