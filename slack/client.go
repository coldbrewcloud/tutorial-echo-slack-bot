package slack

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/coldbrewcloud/tutorial-echo-slack-bot/utils"
	"golang.org/x/net/websocket"
)

type Client struct {
	apiToken string
}

func NewClient(apiToken string) *Client {
	return &Client{apiToken: apiToken}
}

func (c *Client) Start(messageCallback MessageCallback) error {
	res := &RTMStartResponse{}
	err := c.executeWebAPI("rtm.start", nil, res)
	if err != nil {
		return err
	}

	if !res.OK {
		return fmt.Errorf("Slack error: %s", res.Error)
	}

	wsURL := res.URL
	botID := res.Self.ID
	if wsURL == "" || botID == "" {
		return fmt.Errorf("Invalid API response format: %s", utils.ToJSON(res))
	}

	ws, err := websocket.Dial(wsURL, "", slackWebSocketOrigin)
	if err != nil {
		return err
	}

	for {
		message := &RTMMessage{}
		if err := websocket.JSON.Receive(ws, &message); err != nil {
			return err
		}

		if message.Type != "message" {
			continue
		}

		messagePrefix := fmt.Sprintf("<@%s>", botID)
		if strings.HasPrefix(message.Text, messagePrefix) {
			trimmedMessage := strings.TrimSpace(message.Text[len(messagePrefix):])
			if err := messageCallback(trimmedMessage, message.Channel); err != nil {
				return err
			}
		}
	}

	return nil
}

func (c *Client) PostMessage(message, channel string) error {
	params := map[string]string{
		"channel": channel,
		"text":    message,
	}

	apiResponse := &APIResponse{}
	if err := c.executeWebAPI("chat.postMessage", params, apiResponse); err != nil {
		return err
	}
	if !apiResponse.OK {
		return fmt.Errorf("Slack error: %s", apiResponse.Error)
	}

	return nil
}

func (c *Client) executeWebAPI(apiMethod string, args map[string]string, apiResponse interface{}) error {
	postArgs := url.Values{
		"token": {c.apiToken},
	}
	if args != nil {
		for ak, av := range args {
			postArgs.Set(ak, av)
		}
	}

	res, err := http.PostForm(fmt.Sprintf("%s%s", slackAPIBaseURL, apiMethod), postArgs)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return fmt.Errorf("Slack API error response [%d]: %s", res.StatusCode, string(resBody))
	}

	return json.Unmarshal(resBody, apiResponse)
}
