-- name: GetCourts :many
SELECT * FROM court;

-- name: GetCourt :one
SELECT * FROM court WHERE uuid = $1;