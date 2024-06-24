package model

type Members []Member

type Member struct {
	Id         int    `json:"id"`
	Uuid       string `json:"uuid"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Title      string `json:"title"`
	Dob        string `json:"dob"`
	Email      string `json:"email"`
	ProfileUrl string `json:"profile_url"`
	CreatedAt  string `json:"created_at"`
}
