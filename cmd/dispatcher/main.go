package main

import (
	"fmt"
	"github.com/slack-go/slack"
	"io/ioutil"
	"log"
	"os"
)

//type messages


func main()  {

	file, err := os.Open("../../data/example.json");
	if err != nil {
		log.Fatal(err)
	}

	bytes, err := ioutil.ReadAll(file)



	api := slack.New("")

	attachment := slack.Attachment{
		Pretext: "some pretext",
		Text:    "some text",
	}

	channelID, timestamp, err := api.PostMessage(
		"C02H5DPC91S",
		slack.MsgOptionText("Some text", false),
		slack.MsgOptionAttachments(attachment),
		slack.MsgOptionAsUser(false), // Add this if you want that the bot would post message as a user, otherwise it will send response using the default slackbot
	)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)

}