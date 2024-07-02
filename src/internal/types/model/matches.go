package model

import (
	"github.com/google/uuid"
)

type TeamScore struct {
	Games   []int  `json:"games"`
	Current string `json:"current"`
}

type MatchTeamsScore struct {
	A TeamScore `json:"a"`
	B TeamScore `json:"b"`
}

type MatchScore struct {
	Uuid  uuid.UUID       `json:"uuid"`
	Score MatchTeamsScore `json:"score"`
}

type Match struct {
	Uuid  uuid.UUID       `json:"uuid"`
	Score MatchTeamsScore `json:"score"`
}

type Matches []Match
