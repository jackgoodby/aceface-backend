-- name: GetMatches :many
SELECT * FROM match;

-- name: GetMatchScore :one
SELECT uuid, score FROM match WHERE uuid = $1;

-- name: GetMatch :one
SELECT * FROM match WHERE uuid = $1;

-- name: UpdateMatchScore :exec
UPDATE match SET score = $2 WHERE uuid = $1;