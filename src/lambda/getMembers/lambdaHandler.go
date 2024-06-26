package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/jackgoodby/aceface-backend/internal/adapters"
)

//func lambdaAdapter(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
//
//	members, err := logic()
//	if err != nil {
//		return events.APIGatewayProxyResponse{}, err
//	}
//
//	jsonBody, marshalErr := json.Marshal(members)
//	if marshalErr != nil {
//		return events.APIGatewayProxyResponse{}, marshalErr
//	}
//
//	return events.APIGatewayProxyResponse{
//		StatusCode: 200,
//		Body:       string(jsonBody),
//	}, err
//}

// Function to retrieve parameter value from SSM Parameter Store
func getParameter(ctx context.Context, client *ssm.Client, name string) (string, error) {
	input := &ssm.GetParameterInput{
		Name:           aws.String(name),
		WithDecryption: aws.Bool(true),
	}

	result, err := client.GetParameter(ctx, input)
	if err != nil {
		return "", fmt.Errorf("failed to get parameter: %w", err)
	}

	return *result.Parameter.Value, nil
}

func HandleRequest(ctx context.Context) (string, error) {
	ssmClient := adapters.GetSsmClient()
	paramName := "LAMBDA_ENV"
	paramValue, err := getParameter(ctx, ssmClient, paramName)

	if err != nil {
		return "", fmt.Errorf("failed to get parameter named %v. Did you set it in parameter store?", err)
	}

	return fmt.Sprintf("LAMBDA_ENV: %s", paramValue), nil
}

func main() {
	lambda.Start(HandleRequest)
}
