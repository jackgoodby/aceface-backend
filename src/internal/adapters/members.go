package adapters

import (
	"github.com/jackgoodby/aceface-backend/internal/store"
	"github.com/jackgoodby/aceface-backend/internal/types/model"
)

// For handling members

type DdbMemberItem struct {
	Id          string `dynamodbav:"id"`
	SecondaryId string `dynamodbav:"secondaryId"`
	store.Member
}

func ReadMembers(applicationId string) (*DdbMemberItem, error) {
	key := &KeyBasedStruct{
		Id:          applicationId,
		SecondaryId: "application",
	}

	result := &DdbMemberItem{}
	_, getItemErr := dynamodbGetWrapper(key, result)
	if getItemErr != nil {
		return &DdbMemberItem{}, getItemErr
	}

	return result, nil
}

func ReadMember(memberId int) (*model.Member, error) {
	result := &model.Member{
		Id:         1,
		Uuid:       "12345-234234-234234-23423",
		FirstName:  "moofirst",
		LastName:   "mooLast",
		Title:      "Mr. ",
		Dob:        "1977-03-08",
		Email:      "moo@moo.com",
		ProfileUrl: "moo",
		CreatedAt:  "2024-06-24 11:26:00",
	}

	return result, nil
}
