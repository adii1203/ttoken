// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: key.sql

package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createKey = `-- name: CreateKey :one
INSERT INTO api_keys (
    project_id, prefix, key_hash, scopes, environment, expires_at
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING id, project_id, prefix, key_hash, scopes, environment, revoked, created_at, expires_at
`

type CreateKeyParams struct {
	ProjectID   uuid.UUID
	Prefix      pgtype.Text
	KeyHash     string
	Scopes      []string
	Environment string
	ExpiresAt   pgtype.Timestamp
}

func (q *Queries) CreateKey(ctx context.Context, arg CreateKeyParams) (ApiKey, error) {
	row := q.db.QueryRow(ctx, createKey,
		arg.ProjectID,
		arg.Prefix,
		arg.KeyHash,
		arg.Scopes,
		arg.Environment,
		arg.ExpiresAt,
	)
	var i ApiKey
	err := row.Scan(
		&i.ID,
		&i.ProjectID,
		&i.Prefix,
		&i.KeyHash,
		&i.Scopes,
		&i.Environment,
		&i.Revoked,
		&i.CreatedAt,
		&i.ExpiresAt,
	)
	return i, err
}
