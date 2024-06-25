package main

import (
	"github.com/google/uuid"
	"github.com/jackgoodby/aceface-backend/internal/adapters"
	"github.com/jackgoodby/aceface-backend/internal/types/model"
)

func logic(uuid uuid.UUID) (*model.Member, error) {
	item, readAppErr := adapters.GetMember(uuid)
	return item, readAppErr
}
