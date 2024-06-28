package main

import (
	"github.com/google/uuid"
	"github.com/jackgoodby/aceface-backend/internal/adapters"
	"github.com/jackgoodby/aceface-backend/internal/types/model"
)

func logic(uuid uuid.UUID) (*model.Court, error) {
	item, readAppErr := adapters.GetCourt(uuid)
	return item, readAppErr
}
