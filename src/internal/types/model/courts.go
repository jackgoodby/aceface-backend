package model

import (
	"github.com/google/uuid"
)

type Courts []Court

type Court struct {
	Uuid        uuid.UUID `json:"uuid"`
	CourtNumber int32     `json:"court_number"`
	AltName     string    `json:"alt_name"`
	Surface     string    `json:"surface"`
}
