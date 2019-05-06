package main

import (
	"fmt"
	"os"

	"github.com/ashwanthkumar/slack-go-webhook"
//	"github.com/k0kubun/pp"
	mkr "github.com/mackerelio/mackerel-client-go"
)

const (
	WEBHOOKURL = "https://hooks.slack.com/xxxxxxxxxxxxxxx"
	CHANNEL    = "xxxxxxxxxxxxx"
	USERNAME   = "mackerel"
)

func main() {
	hostsMap := FindHosts()
	client := mkr.NewClient(os.Getenv("MACKEREL_APIKEY"))
	alerts, err := client.FindAlerts()

//	pp.Println(alerts)
	// エラー処理
	if err != nil {
		os.Exit(1)
	}

	// SprintfでPostSlackを送る
	for _, res := range alerts.Alerts {
		message := fmt.Sprintf(" AlertType: %v\n AlertHostID: %v\n AlertHost: %v\n AlertStatus: %v\n AlertMessage: %v\n", res.Type, res.HostID, hostsMap[res.HostID], res.Status, res.Message)
		PostSlack(message)
	}
}

// FindHosts
func FindHosts() map[string]string {
	client := mkr.NewClient(os.Getenv("MACKEREL_APIKEY"))
	hosts, _ := client.FindHosts(&mkr.FindHostsParam{
		Statuses: []string{"working", "standby", "maintenance", "poweroff"},
	})

	// Hostsを連想配列に
	var m map[string]string
	m = make(map[string]string)

	for _, v := range hosts {
//		fmt.Println(v.Name, v.ID)
		m[v.ID] = v.Name

	}
//	pp.Println(hosts)
	return m
}

// Slack通知
func PostSlack(msg string) {
	field1 := slack.Field{Title: "Message", Value: msg}

	attachment := slack.Attachment{}
	attachment.AddField(field1)

	color := "good"
	attachment.Color = &color
	payload := slack.Payload{
		Username:    USERNAME,
		Channel:     CHANNEL,
		Attachments: []slack.Attachment{attachment},
	}
	err := slack.Send(WEBHOOKURL, "", payload)
	if err != nil {
		os.Exit(1)
	}
}

