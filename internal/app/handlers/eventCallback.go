package handlers

import (
	"encoding/json"

	lambdaEvents "github.com/aws/aws-lambda-go/events"
	slackEvents "github.com/smort/quoteboard_slack/internal/pkg/slack/events"
)

func HandleEventCallback(event *slackEvents.EventWrapper) (lambdaEvents.APIGatewayProxyResponse, error) {
	var generic slackEvents.EventGeneric
	json.Unmarshal([]byte(event.Event), &generic)

	if generic.Type == slackEvents.MessageType {
		HandleMessage(event.Event)
	}

	return lambdaEvents.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}
