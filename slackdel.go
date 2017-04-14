package main

import (
	"log"

	"github.com/nlopes/slack"
)

const (
apiToken = "API TOKEN"
channelName = "channel name" // without the octothorpe "#"
)

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalf(msg, err.Error())
	}
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	c := slack.New(apiToken)

	list, err := c.GetChannels(true)
	checkErr(err, "GetChannels: %s")

	var generalChannel slack.Channel
	for _, c := range list {
		if c.Name == channelName {
			generalChannel = c
			break
		}
	}

	hasMore := true
	for hasMore {
		hist, err := c.GetChannelHistory(generalChannel.ID, slack.HistoryParameters{Count: 10000000})
		checkErr(err, "GetChannelHistory: %s")
		hasMore = hist.HasMore
		log.Printf("got %d messages", len(hist.Messages))

		for _, m := range hist.Messages {
			_, _, err = c.DeleteMessage(generalChannel.ID, m.Timestamp)
			checkErr(err, "DeleteMessage: %s")
			log.Printf("deleted %q", m.Text)
		}
	}
}