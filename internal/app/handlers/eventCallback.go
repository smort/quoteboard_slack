package handlers

import (
	"encoding/json"

	lambdaEvents "github.com/aws/aws-lambda-go/events"
	log "github.com/sirupsen/logrus"
	slackEvents "github.com/smort/quoteboard_slack/internal/pkg/slack/events"
)

func HandleEventCallback(event *slackEvents.EventWrapper) (lambdaEvents.APIGatewayProxyResponse, error) {
	log.Debug("Handling event callback")
	var generic slackEvents.EventGeneric
	err := json.Unmarshal([]byte(event.Event), &generic)
	if err != nil {
		log.Errorf("Error while unmarshalling json %v", err)
		return lambdaEvents.APIGatewayProxyResponse{}, err
	}

	if generic.Type == slackEvents.MessageType {
		HandleMessage(event.Event)
	}

	log.Debug("Returning from generic event handler")
	return lambdaEvents.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}
