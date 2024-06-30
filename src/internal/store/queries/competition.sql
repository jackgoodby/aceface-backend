-- name: GetCompetitions :many
SELECT * FROM competition;

-- name: GetCompetition :one
SELECT * FROM competition WHERE uuid = $1;