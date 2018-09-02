package events

// EventWrapper is the event wrapping all other events
type EventWrapper struct {
	Token       string      `json:"token"`
	TeamID      string      `json:"team_id"`
	APIAppID    string      `json:"api_app_id"`
	Event       interface{} `json:"event"`
	EventType   string      `json:"type"`
	AuthedUsers []string    `json:"authed_users"`
	EventID     string      `json:"event_id"`
	EventTime   int64       `json:"event_time"`
}

// ChallengeRequest https://api.slack.com/events-api#prepare
type ChallengeRequest struct {
	Token     string `json:"token"`
	Challenge string `json:"challenge"`
	Type      string `json:"type"`
}

type ChallengeResponse struct {
	ChallengeVal string `json:"challenge"`
}

const (
	ChallengeType = "url_verification"
)
