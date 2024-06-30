package main

import (
	"github.com/google/uuid"
	"github.com/jackgoodby/aceface-backend/internal/adapters"
	"github.com/jackgoodby/aceface-backend/internal/types/model"
)

func logic(uuid uuid.UUID) (*model.Competition, error) {
	item, readAppErr := adapters.GetCompetition(uuid)
	return item, readAppErr
}
