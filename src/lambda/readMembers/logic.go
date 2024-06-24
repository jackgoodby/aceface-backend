package main

import (
	"github.com/jackgoodby/aceface-backend/internal/adapters"
	"github.com/jackgoodby/aceface-backend/internal/types/model"
)

func logic(memberId int) (*model.Member, error) {

	memberItem, readAppErr := adapters.ReadMember(memberId)
	return memberItem, readAppErr
}
