package main

import (
	"os"

	"github.com/coldbrewcloud/tutorial-echo-slack-bot/slack"
)

func main() {
	slackClient := slack.NewClient(os.Getenv("SLACK_API_TOKEN"))
	if err := slackClient.Start(slackClient.PostMessage); err != nil {
		panic(err)
	}
}
