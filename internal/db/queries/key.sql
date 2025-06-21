-- name: CreateKey :one
INSERT INTO api_keys (
    project_id, prefix, key_hash, scopes, environment, expires_at
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING *;