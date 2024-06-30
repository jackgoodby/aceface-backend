package adapters

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackgoodby/aceface-backend/internal/store"
	"github.com/jackgoodby/aceface-backend/internal/types/model"
)

func GetCompetitions() (*model.Competitions, error) {
	ctx := context.Background()

	//logger := zap.Must(zap.NewProduction(zap.Fields(zap.String("service_name", "opg-sirius-finance-api")))).Sugar()

	// get completed string from config - get username and pass from secretsmanager
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
	results, err := storeQueries.GetCompetitions(ctx)
	if err != nil {
		return nil, err
	}

	var competitions model.Competitions

	for _, comp := range results {
		var competition = model.Competition{
			Uuid:       comp.Uuid,
			CompType:   comp.CompType,
			Name:       comp.Name,
			Identifier: comp.Identifier.String,
			IsInternal: comp.IsInternal,
			StartDate:  comp.StartDate,
			EndDate:    comp.EndDate,
		}

		competitions = append(competitions, competition)
	}

	return &competitions, nil
}

func GetCompetition(uuid uuid.UUID) (*model.Competition, error) {
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
	result, err := storeQueries.GetCompetition(ctx, uuid)
	if err != nil {
		return nil, err
	}

	competition := &model.Competition{
		Uuid:       result.Uuid,
		CompType:   result.CompType,
		Name:       result.Name,
		Identifier: result.Identifier.String,
		IsInternal: result.IsInternal,
		StartDate:  result.StartDate,
		EndDate:    result.EndDate,
	}

	return competition, nil
}
