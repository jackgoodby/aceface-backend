package main

import (
	"github.com/jackgoodby/aceface-backend/internal/adapters"
	"github.com/jackgoodby/aceface-backend/internal/types/model"
)

func logic(id int) (*model.Member, error) {
	item, readAppErr := adapters.GetMember(id)
	return item, readAppErr
}
