package model

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackgoodby/aceface-backend/internal/common"
)

type Members []Member

type Member struct {
	Uuid       pgtype.UUID `json:"uuid"`
	FirstName  string      `json:"first_name"`
	LastName   string      `json:"last_name"`
	Title      string      `json:"title"`
	Dob        common.Date `json:"dob"`
	Email      string      `json:"email"`
	ProfileUrl string      `json:"profile_url"`
	CreatedAt  common.Date `json:"created_at"`
}
