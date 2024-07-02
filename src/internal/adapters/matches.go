package adapters

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackgoodby/aceface-backend/internal/store"
	"github.com/jackgoodby/aceface-backend/internal/types/model"
	"reflect"
)

func GetMatches() (*model.Matches, error) {
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
	results, err := storeQueries.GetMatches(ctx)
	if err != nil {
		return nil, err
	}

	var matches model.Matches

	for _, result := range results {
		var match = buildMatchFromResult(result)
		matches = append(matches, match)
	}

	return &matches, nil
}

func GetMatch(uuid uuid.UUID) (*model.Match, error) {
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
	result, err := storeQueries.GetMatch(ctx, uuid)
	if err != nil {
		return nil, err
	}

	match := buildMatchFromResult(result)

	return &match, nil
}

func GetMatchScore(uuid uuid.UUID) (*model.MatchScore, error) {
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
	result, err := storeQueries.GetMatch(ctx, uuid)
	if err != nil {
		return nil, err
	}

	_, matchTeamsScore := validateMatchScore(result.Score)

	matchScore := model.MatchScore{
		Uuid:  result.Uuid,
		Score: matchTeamsScore,
	}

	return &matchScore, nil
}

func UpdateMatchScore(uuid uuid.UUID, matchTeamsScore model.MatchTeamsScore) (*model.MatchScore, error) {
	ctx := context.Background()

	matchScore := &model.MatchScore{
		Uuid:  uuid,
		Score: matchTeamsScore,
	}

	// get completed string from config
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
		config.DBUser,
		config.DBPass,
		config.DBHost,
		config.DBPort,
		config.DBName)

	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		fmt.Printf("Couldn't connect: %v\n", err)
	}
	defer conn.Close(ctx)

	score, err := json.Marshal(matchTeamsScore)
	if err != nil {
		fmt.Printf("Invalid format: %v\n", err)
	}

	storeQueries := store.New(conn)
	updateParams := store.UpdateMatchScoreParams{
		Uuid:  uuid,
		Score: score,
	}
	err = storeQueries.UpdateMatchScore(ctx, updateParams)
	if err != nil {
		fmt.Printf("Update failed: %v\n", err)
	} else {
		matchScore.Score = matchTeamsScore
	}

	return matchScore, nil
}

func buildMatchFromResult(result store.Match) model.Match {
	err, matchTeamsScore := validateMatchScore(result.Score)
	if err != nil {
		fmt.Printf("Score validation failed: %v\n", err)
	} else {
		fmt.Println("Score is valid!")
	}

	match := model.Match{
		Uuid:  result.Uuid,
		Score: matchTeamsScore,
	}

	return match
}

func validateMatchScore(score []byte) (error, model.MatchTeamsScore) {
	var matchTeamsScore model.MatchTeamsScore
	err := json.Unmarshal(score, &matchTeamsScore)
	if err != nil {
		return fmt.Errorf("invalid format: %w", err), matchTeamsScore
	}
	if !isMatchTeamsScore(matchTeamsScore) {
		return fmt.Errorf("score does not match matchTeamsScore struct"), matchTeamsScore
	}
	// Additional checks can be added here if necessary
	return nil, matchTeamsScore
}

func isMatchTeamsScore(score interface{}) bool {
	expectedType := reflect.TypeOf(model.MatchTeamsScore{})
	return reflect.TypeOf(score) == expectedType
}
