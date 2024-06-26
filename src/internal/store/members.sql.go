// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: members.sql

package store

import (
	"context"

	"github.com/google/uuid"
)

const getMember = `-- name: GetMember :one
SELECT id, uuid, first_name, last_name, title, dob, email, profile_url, created_at FROM member WHERE uuid = $1
`

func (q *Queries) GetMember(ctx context.Context, argUuid uuid.UUID) (Member, error) {
	row := q.db.QueryRow(ctx, getMember, argUuid)
	var i Member
	err := row.Scan(
		&i.ID,
		&i.Uuid,
		&i.FirstName,
		&i.LastName,
		&i.Title,
		&i.Dob,
		&i.Email,
		&i.ProfileUrl,
		&i.CreatedAt,
	)
	return i, err
}

const getMembers = `-- name: GetMembers :many
SELECT id, uuid, first_name, last_name, title, dob, email, profile_url, created_at FROM member
`

func (q *Queries) GetMembers(ctx context.Context) ([]Member, error) {
	rows, err := q.db.Query(ctx, getMembers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Member
	for rows.Next() {
		var i Member
		if err := rows.Scan(
			&i.ID,
			&i.Uuid,
			&i.FirstName,
			&i.LastName,
			&i.Title,
			&i.Dob,
			&i.Email,
			&i.ProfileUrl,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
