package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handleEvent(event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println(fmt.Sprintf("Call received: %+v", event))

	// client := slack.New("test")

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "{message:\"test\"}",
	}, nil
}

func main() {
	lambda.Start(handleEvent)
}
