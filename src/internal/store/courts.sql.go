// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: courts.sql

package store

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const getCourt = `-- name: GetCourt :one
SELECT
    c.uuid,
    c.court_number,
    c.alt_name,
    c.surface,
    css.slot_times
FROM court c
         JOIN court_slot_set css
              ON css.uuid = c.slot_set_id
WHERE c.uuid = $1
`

type GetCourtRow struct {
	Uuid        uuid.UUID
	CourtNumber int32
	AltName     pgtype.Text
	Surface     string
	SlotTimes   []byte
}

func (q *Queries) GetCourt(ctx context.Context, argUuid uuid.UUID) (GetCourtRow, error) {
	row := q.db.QueryRow(ctx, getCourt, argUuid)
	var i GetCourtRow
	err := row.Scan(
		&i.Uuid,
		&i.CourtNumber,
		&i.AltName,
		&i.Surface,
		&i.SlotTimes,
	)
	return i, err
}

const getCourts = `-- name: GetCourts :many
SELECT
    c.uuid,
    c.court_number,
    c.alt_name,
    c.surface,
    css.slot_times
FROM court c
JOIN court_slot_set css
    ON css.uuid = c.slot_set_id
`

type GetCourtsRow struct {
	Uuid        uuid.UUID
	CourtNumber int32
	AltName     pgtype.Text
	Surface     string
	SlotTimes   []byte
}

func (q *Queries) GetCourts(ctx context.Context) ([]GetCourtsRow, error) {
	rows, err := q.db.Query(ctx, getCourts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetCourtsRow
	for rows.Next() {
		var i GetCourtsRow
		if err := rows.Scan(
			&i.Uuid,
			&i.CourtNumber,
			&i.AltName,
			&i.Surface,
			&i.SlotTimes,
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
