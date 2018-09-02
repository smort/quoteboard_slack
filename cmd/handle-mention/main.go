package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/nlopes/slack"
	"github.com/nlopes/slack/slackevents"
)

var token = os.Getenv("SLACK_OATH")
var api = slack.New(token)

func handleEvent(event events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	fmt.Println(fmt.Sprintf("Call received: %+v", event))

	slackEvent, err := slackevents.ParseEvent(json.RawMessage(event.Body))
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Headers:    map[string]string{"Content-Type": "application/json"},
			Body:       "{message:\"Error parsing Slack event\"}",
		}
	}

	if slackEvent.Type == slackevents.URLVerification {
		var r *slackevents.ChallengeResponse
		err := json.Unmarshal([]byte(event.Body), &r)
		if err != nil {
			return events.APIGatewayProxyResponse{
				StatusCode: 500,
				Headers:    map[string]string{"Content-Type": "text/plain"},
				Body:       r.Challenge,
			}
		}

		return events.APIGatewayProxyResponse{
			StatusCode: 200,
		}
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       "{message:\"test\"}",
	}
}

func main() {
	lambda.Start(handleEvent)
}
