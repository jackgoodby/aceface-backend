package main

import (
	"github.com/google/uuid"
	"github.com/jackgoodby/aceface-backend/internal/adapters"
	"github.com/jackgoodby/aceface-backend/internal/types/model"
)

func logic(uuid uuid.UUID, updates model.MatchTeamsScore) (*model.MatchScore, error) {

	// add a LOT of logic around what represents a valid score.
	// current 0,15,30,40,A - Cannot be A + A. Cannot be A + !40. 0 should be empty string
	// games: complete set criteria: 6 + <= 4 or 7,5 or 7,6. Not 7 + <5. Not 7 + 7
	// games: If 2nd array element in use (2nd set in play) then 1st element must meet completed set criteria above
	// If set is complete then begin new set 0,0
	// If set 1 is made incomplete (by submitting incomplete score), delete 2nd set if it is 0,0
	// valid in play: [0-5],[0-5] or 6,[5-6] (or reverse)

	item, updateErr := adapters.UpdateMatchScore(uuid, updates)
	return item, updateErr
}
