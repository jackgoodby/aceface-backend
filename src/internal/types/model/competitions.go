package model

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Competitions []Competition

type Competition struct {
	Uuid       uuid.UUID          `json:"uuid"`
	CompType   string             `json:"comp_type"`
	Name       string             `json:"name"`
	Identifier string             `json:"identifier"`
	IsInternal bool               `json:"is_internal"`
	StartDate  pgtype.Timestamptz `json:"start_date"`
	EndDate    pgtype.Timestamptz `json:"end_date"`
}
