package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/google/uuid"
	"github.com/jackgoodby/aceface-backend/internal/common"
	"github.com/jackgoodby/aceface-backend/internal/types"
	"github.com/jackgoodby/aceface-backend/internal/types/model"
	"net/http"
)

func lambdaAdapter(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	id, err := uuid.Parse(request.PathParameters["uuid"])
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       fmt.Sprint("Invalid path parameter: uuid must be a UUID"),
		}, nil
	}

	var validBody model.MatchTeamsScore
	unmarshalErr := json.Unmarshal([]byte(request.Body), &validBody)
	if unmarshalErr != nil {
		panic(unmarshalErr)
	}

	matchScore, err := logic(id, validBody)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	jsonBody, marshalErr := json.Marshal(matchScore)
	if marshalErr != nil {
		return events.APIGatewayProxyResponse{}, marshalErr
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(jsonBody),
	}, err
}

func getLambdaHandler() types.HandlerSignature {
	wrappedLambdaAdapter := common.LambdaWrapper(lambdaAdapter)
	return wrappedLambdaAdapter
}

func main() {
	lambda.Start(getLambdaHandler())
}
