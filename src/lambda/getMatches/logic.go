package main

import (
	"github.com/jackgoodby/aceface-backend/internal/adapters"
	"github.com/jackgoodby/aceface-backend/internal/types/model"
)

func logic() (*model.Matches, error) {
	items, readAppErr := adapters.GetMatches()
	return items, readAppErr
}
