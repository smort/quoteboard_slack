package main

import (
	"encoding/json"

	lambdaEvents "github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	log "github.com/sirupsen/logrus"
	"github.com/smort/quoteboard_slack/internal/app/handlers"
	slackEvents "github.com/smort/quoteboard_slack/internal/pkg/slack/events"
	"github.com/smort/quoteboard_slack/internal/pkg/utils"
)

// var token = os.Getenv("SLACK_OATH")
// var api = slack.New(token)

var hasInited = false

func init() {
	log.Debugf("Init called. Has already init'd? %v", hasInited)
	if hasInited == true {
		return
	}

	log.SetLevel(log.DebugLevel)
	log.Debug("LogLevel set")

	hasInited = true
}

func handleEvent(event lambdaEvents.APIGatewayProxyRequest) (lambdaEvents.APIGatewayProxyResponse, error) {
	log.Debugf("Event received %v", event)

	var slackEvent *slackEvents.EventWrapper
	err := json.Unmarshal([]byte(event.Body), &slackEvent)
	if err != nil {
		log.Errorf("Error while parsing event body. %v", err)
		return utils.RespondWithError(err), nil
	}

	var handlerError error
	var response lambdaEvents.APIGatewayProxyResponse
	switch slackEvent.EventType {
	case slackEvents.ChallengeType:
		log.Debug("Event is a challenge")
		response, handlerError = handlers.HandleChallenge(event.Body)
	case slackEvents.CallbackType:
		log.Debug("Event is an event callback")
		response, handlerError = handlers.HandleEventCallback(slackEvent)
	default:
		log.Debug("Event is unknown type, responding with default")
		response = lambdaEvents.APIGatewayProxyResponse{
			StatusCode: 200,
		}
	}

	if handlerError != nil {
		log.Errorf("An error occurred in a handler. %v", handlerError)
		// returning 500 will cause Slack to retry
		return lambdaEvents.APIGatewayProxyResponse{
			StatusCode: 500,
		}, nil
	} else {
		return response, nil
	}
}

func main() {
	lambda.Start(handleEvent)
}
