-- name: CreateUser :one
INSERT INTO users (
    clerk_id,
    email,
    first_name,
    last_name
) VALUES (
    $1,
    $2,
    $3,
    $4
) RETURNING *;