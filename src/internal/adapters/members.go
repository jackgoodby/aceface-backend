package adapters

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackgoodby/aceface-backend/internal/common"
	"github.com/jackgoodby/aceface-backend/internal/store"
	"github.com/jackgoodby/aceface-backend/internal/types/model"
)

func stringToPgtypeText(s string) pgtype.Text {
	var pgText pgtype.Text
	pgText.String = s
	//pgText.Status = pgtype.Present
	return pgText
}

func GetMembers() (*model.Members, error) {
	ctx := context.Background()

	//logger := zap.Must(zap.NewProduction(zap.Fields(zap.String("service_name", "opg-sirius-finance-api")))).Sugar()

	// get completed string from config
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
		config.DBUser,
		config.DBPass,
		config.DBHost,
		config.DBPort,
		config.DBName)

	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		return nil, err
	}
	defer conn.Close(ctx)

	storeQueries := store.New(conn)
	results, err := storeQueries.GetMembers(ctx)
	if err != nil {
		return nil, err
	}

	var members model.Members

	for _, mem := range results {
		var member = model.Member{
			Uuid:       mem.Uuid,
			FirstName:  mem.FirstName.String,
			LastName:   mem.LastName.String,
			Title:      mem.Title.String,
			Dob:        common.Date{Time: mem.Dob.Time},
			Email:      mem.Email.String,
			ProfileUrl: mem.ProfileUrl.String,
			CreatedAt:  common.Date{Time: mem.CreatedAt.Time},
		}

		members = append(members, member)
	}

	return &members, nil
}

func GetMember(uuid uuid.UUID) (*model.Member, error) {
	ctx := context.Background()

	// get completed string from config
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
		config.DBUser,
		config.DBPass,
		config.DBHost,
		config.DBPort,
		config.DBName)

	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		return nil, err
	}
	defer conn.Close(ctx)

	//convert uuid into a pgUUID
	var pgUUID pgtype.UUID
	copy(pgUUID.Bytes[:], uuid[:])
	//pgUUID.Status = pgtype.Present

	storeQueries := store.New(conn)
	result, err := storeQueries.GetMember(ctx, pgUUID)
	if err != nil {
		return nil, err
	}

	member := &model.Member{
		Uuid:       result.Uuid,
		FirstName:  result.FirstName.String,
		LastName:   result.LastName.String,
		Title:      result.Title.String,
		Dob:        common.Date{Time: result.Dob.Time},
		Email:      result.Email.String,
		ProfileUrl: result.ProfileUrl.String,
		CreatedAt:  common.Date{Time: result.CreatedAt.Time},
	}

	return member, nil
}
