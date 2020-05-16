package slacknotify

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

// SlackNotifierConfig contains the configuration for a SlackNotifier object
type SlackNotifierConfig struct {
	WebhookURL     string
	RequestTimeout time.Duration
}

var defaultConf SlackNotifierConfig = SlackNotifierConfig{
	WebhookURL:     "https://hooks.slack.com/services/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",
	RequestTimeout: 10 * time.Second,
}

// SlackNotifier used to send notifications to slack
type SlackNotifier struct {
	config SlackNotifierConfig
}

type slackNotification struct {
	Text string `json:"text"`
}

// Init creates a new SlackNotifier object
func Init(conf *SlackNotifierConfig) (*SlackNotifier, error) {
	if conf == nil {
		return &SlackNotifier{defaultConf}, nil
	}

	if conf.WebhookURL == "" {
		conf.WebhookURL = defaultConf.WebhookURL
	}

	if conf.RequestTimeout.Seconds() == 0 {
		conf.RequestTimeout = defaultConf.RequestTimeout
	}

	return &SlackNotifier{
		config: *conf,
	}, nil
}

// SendString uses the SlackNotifier to send a simple text message
func (sn *SlackNotifier) SendString(message string) error {
	body, err := json.Marshal(slackNotification{message})
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", sn.config.WebhookURL, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{Timeout: sn.config.RequestTimeout}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	buf := &bytes.Buffer{}
	_, err = buf.ReadFrom(resp.Body)
	if buf.String() != "ok" {
		return err
	}

	return nil
}

func (sn *SlackNotifier) Write(b []byte) (int, error) {
	return 0, nil
}
