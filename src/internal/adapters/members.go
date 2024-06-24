package adapters

import (
	"fmt"
	"github.com/jackgoodby/aceface-backend/internal/store"
	"github.com/jackgoodby/aceface-backend/internal/types/model"
)

// For handling members

type DdbMemberItem struct {
	Id          string `dynamodbav:"id"`
	SecondaryId string `dynamodbav:"secondaryId"`
	store.Member
}

//func GetMembers() (*model.Members, error) {
//
//	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
//		os.Getenv("DB_HOST"),
//		os.Getenv("DB_PORT"),
//		os.Getenv("DB_USER"),
//		os.Getenv("DB_PASSWORD"),
//		os.Getenv("DB_NAME"))
//
//	db, err := sql.Open("postgres", connStr)
//	if err != nil {
//		log.Fatalf("Error opening database: %q", err)
//		return events.APIGatewayProxyResponse{StatusCode: 500}, err
//	}
//	defer db.Close()
//
//	rows, err := db.Query("SELECT id, uuid, first_name, last_name, title, dob, profile_url, created_at, somebool FROM test_person")
//	if err != nil {
//		log.Fatalf("Error fetching data: %q", err)
//		return events.APIGatewayProxyResponse{StatusCode: 500}, err
//	}
//	defer rows.Close()
//
//	var persons []Person
//	for rows.Next() {
//		var p Person
//		if err := rows.Scan(&p.ID, &p.UUID, &p.FirstName, &p.LastName, &p.Title, &p.DOB, &p.ProfileURL, &p.CreatedAt, &p.SomeBool); err != nil {
//			log.Fatalf("Error scanning row: %q", err)
//			return events.APIGatewayProxyResponse{StatusCode: 500}, err
//		}
//		persons = append(persons, p)
//	}
//
//	return result, nil
//}

func GetMember(memberId int) (*model.Member, error) {

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost,
		config.DBPort,
		config.DBUser,
		config.DBPass,
		config.DBName)

	result := &model.Member{
		Id:         1,
		Uuid:       connStr,
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
