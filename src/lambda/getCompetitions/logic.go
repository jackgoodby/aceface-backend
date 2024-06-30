package main

import (
	"github.com/jackgoodby/aceface-backend/internal/adapters"
	"github.com/jackgoodby/aceface-backend/internal/types/model"
)

func logic() (*model.Competitions, error) {
	items, readAppErr := adapters.GetCompetitions()
	return items, readAppErr
}
