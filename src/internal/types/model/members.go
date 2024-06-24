package model

import "github.com/jackgoodby/aceface-backend/internal/common"

type Members []Member

type Member struct {
	Id         int         `json:"id"`
	Uuid       string      `json:"uuid"`
	FirstName  string      `json:"first_name"`
	LastName   string      `json:"last_name"`
	Title      string      `json:"title"`
	Dob        common.Date `json:"dob"`
	Email      string      `json:"email"`
	ProfileUrl string      `json:"profile_url"`
	CreatedAt  common.Date `json:"created_at"`
}
