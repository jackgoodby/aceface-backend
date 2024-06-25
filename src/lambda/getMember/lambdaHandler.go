package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jackgoodby/aceface-backend/internal/common"
	"github.com/jackgoodby/aceface-backend/internal/types"
	"net/http"
	"strconv"
)

func lambdaAdapter(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	s := request.PathParameters["id"]

	id, err := strconv.Atoi(s)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       fmt.Sprintf("Invalid path parameter: %s must be an integer", s),
		}, nil
	}

	// Accessing query string parameters
	//queryParam := request.QueryStringParameters["yourQueryParameter"]
	//fmt.Printf("Query Parameter: %s\n", queryParam)

	member, err := logic(id)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	jsonBody, marshalErr := json.Marshal(member)
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
