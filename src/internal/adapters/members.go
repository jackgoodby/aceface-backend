package adapters

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackgoodby/aceface-backend/internal/common"
	"github.com/jackgoodby/aceface-backend/internal/store"
	"github.com/jackgoodby/aceface-backend/internal/types/model"
)

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
			FirstName:  mem.FirstName,
			LastName:   mem.LastName,
			Title:      mem.Title,
			Dob:        common.Date{Time: mem.Dob.Time},
			Email:      mem.Email,
			ProfileUrl: mem.ProfileUrl,
			CreatedAt:  common.Date{Time: mem.CreatedAt.Time},
		}

		members = append(members, member)
	}

	return &members, nil
}

func GetMember(memberId int) (*model.Member, error) {
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

	storeQueries := store.New(conn)
	result, err := storeQueries.GetMember(ctx, int32(memberId))
	if err != nil {
		return nil, err
	}

	member := &model.Member{
		Uuid:       result.Uuid,
		FirstName:  result.FirstName,
		LastName:   result.LastName,
		Title:      result.Title,
		Dob:        common.Date{Time: result.Dob.Time},
		Email:      result.Email,
		ProfileUrl: result.ProfileUrl,
		CreatedAt:  common.Date{Time: result.CreatedAt.Time},
	}

	return member, nil
}
