package main

import (
	"encoding/json"
	"fmt"

	lambdaEvents "github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/smort/quoteboard_slack/internal/app/handlers"
	slackEvents "github.com/smort/quoteboard_slack/internal/pkg/slack/events"
	"github.com/smort/quoteboard_slack/internal/pkg/utils"
)

// var token = os.Getenv("SLACK_OATH")
// var api = slack.New(token)

func handleEvent(event lambdaEvents.APIGatewayProxyRequest) (lambdaEvents.APIGatewayProxyResponse, error) {
	fmt.Println(fmt.Sprintf("Call received: %+v", event))

	var slackEvent *slackEvents.EventWrapper
	err := json.Unmarshal([]byte(event.Body), &slackEvent)
	if err != nil {
		return utils.RespondWithError(err), nil
	}

	if slackEvent.EventType == slackEvents.ChallengeType {
		return handlers.HandleChallenge(event.Body)
	}

	return lambdaEvents.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handleEvent)
}
