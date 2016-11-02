package slack

const (
	slackAPIBaseURL      = "https://slack.com/api/"
	slackWebSocketOrigin = "https://api.slack.com/"
)

type MessageCallback func(message, channel string) error
