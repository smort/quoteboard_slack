package utils

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

// RespondWithJSON returns APIGatewayProxyResponse that returns 200 with JSON
// if the input can be marshaled to JSON and 500 if not
func RespondWithJSON(body interface{}) events.APIGatewayProxyResponse {
	jsonBody, err := json.Marshal(body)

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Headers:    contentTypeJSON(),
			Body:       fmt.Sprintf("{ message: \"%s\" }", err),
		}
	} else {
		return events.APIGatewayProxyResponse{
			StatusCode: 200,
			Headers:    contentTypeJSON(),
			Body:       string(jsonBody),
		}
	}
}

func RespondWithError(err error) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: 500,
		Headers:    contentTypeJSON(),
		Body:       fmt.Sprintf("{ \"message\": \"%v\" }", err),
	}
}

func contentTypeJSON() map[string]string {
	return map[string]string{
		"Content-Type": "application/json",
	}
}
