package main

import (
	"fmt"
	"github.com/miekg/dns"
	"github.com/slack-go/slack"
	"log"
	"os"
	"time"
)

func main() {
	token := os.Getenv("SLACK_BOT_TOKEN")
	channelId := os.Getenv("CHANNEL_ID")

	dnsServer := os.Getenv("DNS_SERVER")
	domain := os.Getenv("DOMAIN")

	client := slack.New(token)

	ticker := time.NewTicker(time.Millisecond * 1000 * 60 * 60)
	defer ticker.Stop()

	log.Printf("DNS Watchdog task has been started")

	count := 0
	execLookup(*client, channelId, dnsServer, domain, &count)
	for {
		select {
		case <-ticker.C:
			log.Printf("count=%d\n", count)
			execLookup(*client, channelId, dnsServer, domain, &count)
		}
	}
}

func lookupMXRecords(dnsServer string, domain string) bool {
	m := new(dns.Msg)
	m.SetQuestion(domain+".", dns.TypeMX)

	c := new(dns.Client)
	in, _, err := c.Exchange(m, dnsServer)
	if err != nil {
		fmt.Println("DNS Query Failed:", err)
		return false
	}

	for _, answer := range in.Answer {
		if mx, ok := answer.(*dns.MX); ok {
			fmt.Printf("Host: %s, Priority: %d\n", mx.Mx, mx.Preference)
		}
	}
	return true
}

func execLookup(client slack.Client, channelId string, dnsServer string, domain string, count *int) {
	ok := lookupMXRecords(dnsServer, domain)
	if ok {
		*count++
		if *count > 24 {
			*count = 0
			SendDailyNotification(client, channelId, dnsServer, domain)
		}
	} else {
		*count = 0
		SendHourlyNotification(client, channelId, dnsServer, domain)
	}
}
