package handlers

import (
	"encoding/json"

	lambdaEvents "github.com/aws/aws-lambda-go/events"
	slackEvents "github.com/smort/quoteboard_slack/internal/pkg/slack/events"
	"github.com/smort/quoteboard_slack/internal/pkg/utils"
)

func HandleChallenge(body string) (lambdaEvents.APIGatewayProxyResponse, error) {
	var challengeRequest *slackEvents.ChallengeRequest
	err := json.Unmarshal([]byte(body), &challengeRequest)
	if err != nil {
		return utils.RespondWithError(err), nil
	}

	return utils.RespondWithJSON(slackEvents.ChallengeResponse{
		ChallengeVal: challengeRequest.Challenge,
	}), nil
}
