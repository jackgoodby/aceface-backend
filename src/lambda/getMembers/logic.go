package main

import (
	"github.com/jackgoodby/aceface-backend/internal/adapters"
	"github.com/jackgoodby/aceface-backend/internal/types/model"
)

func logic() (*model.Members, error) {
	items, readAppErr := adapters.GetMembers()
	return items, readAppErr
}
