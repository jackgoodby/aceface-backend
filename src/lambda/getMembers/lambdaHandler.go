package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jackgoodby/aceface-backend/internal/common"
	"github.com/jackgoodby/aceface-backend/internal/types"
)

func lambdaAdapter(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	//memberId := request.PathParameters["memberId"]

	memberId := 5
	member, err := logic(memberId)
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