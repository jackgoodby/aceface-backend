-- name: GetMembers :many
SELECT * FROM member;

-- name: GetMember :one
SELECT * FROM member WHERE uuid = $1;