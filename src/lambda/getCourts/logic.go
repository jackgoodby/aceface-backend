package main

import (
	"github.com/jackgoodby/aceface-backend/internal/adapters"
	"github.com/jackgoodby/aceface-backend/internal/types/model"
)

func logic() (*model.Courts, error) {
	items, readAppErr := adapters.GetCourts()
	return items, readAppErr
}
