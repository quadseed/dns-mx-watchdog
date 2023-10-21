package main

import (
	"fmt"
	"github.com/slack-go/slack"
	"log"
	"time"
)

func SendDailyNotification(client slack.Client, channelId string, dnsServer string, domain string) {
	attachment := slack.Attachment{
		Title: "DNS MX Query Successful (daily report)",
		Color: "good",

		Fields: []slack.AttachmentField{
			{
				Title: "Domain",
				Value: "*" + domain + "*",
				Short: false,
			}, {
				Title: "DNS Server",
				Value: "*" + dnsServer + "*",
				Short: false,
			},
			{
				Title: "Date & Time",
				Value: time.Now().Format("2006/1/2 15:04"),
				Short: false,
			},
		},
	}

	_, _, err := client.PostMessage(
		channelId,
		slack.MsgOptionAttachments(attachment),
		slack.MsgOptionAsUser(true),
	)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	log.Printf("[daily report] Message successfully sent to channel")
}

func SendHourlyNotification(client slack.Client, channelId string, dnsServer string, domain string) {
	attachment := slack.Attachment{
		Pretext: "<!channel>",
		Title:   "DNS MX Query FAILED (hourly report)",
		Color:   "danger",

		Fields: []slack.AttachmentField{
			{
				Title: "Domain",
				Value: "*" + domain + "*",
				Short: false,
			}, {
				Title: "DNS Server",
				Value: "*" + dnsServer + "*",
				Short: false,
			},
			{
				Title: "Date & Time",
				Value: time.Now().Format("2006/1/2 15:04"),
				Short: false,
			},
		},
	}

	_, _, err := client.PostMessage(
		channelId,
		slack.MsgOptionAttachments(attachment),
		slack.MsgOptionAsUser(true),
	)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("[hourly report] Message successfully sent to channel")
}
