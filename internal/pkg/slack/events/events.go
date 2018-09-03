package events

import (
	"encoding/json"
)

// EventWrapper is the event wrapping all other events
type EventWrapper struct {
	Token       string          `json:"token"`
	TeamID      string          `json:"team_id"`
	APIAppID    string          `json:"api_app_id"`
	Event       json.RawMessage `json:"event"`
	EventType   string          `json:"type"`
	AuthedUsers []string        `json:"authed_users"`
	EventID     string          `json:"event_id"`
	EventTime   int64           `json:"event_time"`
}

type EventGeneric struct {
	Type string `json:"type"`
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

type MessageEvent struct {
	Type    string `json:"type"`
	Channel string `json:"channel"`
	User    string `json:"user"`
	Text    string `json:"text"`
	TS      int64  `json:"ts"`
}

const (
	CallbackType  = "event_callback"
	ChallengeType = "url_verification"
	MessageType   = "message"
)
