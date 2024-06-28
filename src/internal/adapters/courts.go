package adapters

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackgoodby/aceface-backend/internal/store"
	"github.com/jackgoodby/aceface-backend/internal/types/model"
)

func GetCourts() (*model.Courts, error) {
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
	results, err := storeQueries.GetCourts(ctx)
	if err != nil {
		return nil, err
	}

	var courts model.Courts

	for _, ct := range results {
		var court = model.Court{
			Uuid:        ct.Uuid,
			CourtNumber: ct.CourtNumber,
			AltName:     ct.AltName.String,
			Surface:     ct.Surface,
		}

		courts = append(courts, court)
	}

	return &courts, nil
}

func GetCourt(uuid uuid.UUID) (*model.Court, error) {
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
	result, err := storeQueries.GetCourt(ctx, uuid)
	if err != nil {
		return nil, err
	}

	court := &model.Court{
		Uuid:        result.Uuid,
		CourtNumber: result.CourtNumber,
		AltName:     result.AltName.String,
		Surface:     result.Surface,
	}

	return court, nil
}
