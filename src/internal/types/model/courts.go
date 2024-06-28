package model

import (
	"github.com/google/uuid"
)

type CourtSlotTime struct {
	Duration  int32  `json:"duration"`
	StartTime string `json:"start_time"`
}

type CourtSlotTimes []CourtSlotTime

type Courts []Court

type Court struct {
	Uuid        uuid.UUID      `json:"uuid"`
	CourtNumber int32          `json:"court_number"`
	AltName     string         `json:"alt_name"`
	Surface     string         `json:"surface"`
	SlotTimes   CourtSlotTimes `json:"slot_times"`
}
