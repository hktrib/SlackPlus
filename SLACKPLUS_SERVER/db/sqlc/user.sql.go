// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: user.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT into users (
    username, 
    hashed_password, 
    email
) VALUES ( $1, $2, $3
) RETURNING username, hashed_password, email, password_changed_at, created_at, is_email_verified
`

type CreateUserParams struct {
	Username       string `json:"username"`
	HashedPassword string `json:"hashed_password"`
	Email          string `json:"email"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Username, arg.HashedPassword, arg.Email)
	var i User
	err := row.Scan(
		&i.Username,
		&i.HashedPassword,
		&i.Email,
		&i.PasswordChangedAt,
		&i.CreatedAt,
		&i.IsEmailVerified,
	)
	return i, err
}
