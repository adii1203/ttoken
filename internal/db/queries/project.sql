-- name: CreateProject :one
INSERT INTO projects (name) VALUES ($1) RETURNING *;

-- name: GetProject :one
SELECT * FROM projects WHERE id = $1;