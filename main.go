package main // import "github.com/TailorBrands/slack-cli"

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

type SlackMsg struct {
	Channel   string `json:"channel"`
	Username  string `json:"username,omitempty"`
	Text      string `json:"text"`
	IconEmoji string `json:"icon_emoji,omitempty"`
}

func (m SlackMsg) Encode() (string, error) {
	b, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (m SlackMsg) Post() error {
	encoded, err := m.Encode()
	if err != nil {
		return err
	}

	resp, err := http.PostForm(WebhookUrl, url.Values{"payload": {encoded}})
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("Status code not OK")
	}

	return nil
}

var WebhookUrl string
var Msg SlackMsg

func main() {
	var rootCmd = &cobra.Command{
		Use:   "slack-cli [flags] [string to send]",
		Short: "Send any message to any slack channel",
		Long: `slack-cli is for sending a slack message to any channel
           with configurable flags`,
		RunE: func(cmd *cobra.Command, args []string) error {

			if WebhookUrl == "" {
				if val, ok := os.LookupEnv("WEBHOOK_URL"); !ok {
					return errors.New("Webhook URL not set")
				} else {
					WebhookUrl = val
				}
			}

			Msg.Text = strings.Join(args, " ")

			return Msg.Post()
		},
	}

	rootCmd.PersistentFlags().StringVar(&WebhookUrl, "webhook-url", "", "slack webhook url")
	rootCmd.PersistentFlags().StringVarP(&Msg.Channel, "channel", "c", "alerts", "slack channel to send the message to")
	rootCmd.PersistentFlags().StringVarP(&Msg.Username, "username", "u", "slack-cli", "slack username that sends out the message")
	rootCmd.PersistentFlags().StringVarP(&Msg.IconEmoji, "icon-emoji", "i", "", "user's icon emoji")

	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
