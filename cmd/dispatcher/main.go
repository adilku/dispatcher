package main

import (
	"encoding/json"
	"github.com/slack-go/slack"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

//type messages

type channel struct {
	Text 	 string `json:"text"`
	ChanName string `json:"channel"`
}

type messages struct {
	BotToken string `json:"bot_token"`
	Channels  []channel `json:"channels"`
}

func sendToChan(wg *sync.WaitGroup, mutex *sync.Mutex, chanName string, text string, api *slack.Client) {
	defer wg.Done()
	mutex.Lock()
	log.Println("starting send")
	_, _, err := api.PostMessage(chanName, slack.MsgOptionText(text, false))
	if err != nil {
		log.Println("cant send message", err)
	}
	log.Println("send finished")
	mutex.Unlock()
}

func main()  {
	file, err := os.Open(os.Args[1]);
	if err != nil {
		log.Fatal(err)
	}
	bytes, err := ioutil.ReadAll(file)
	mess := messages{}
	err = json.Unmarshal(bytes, &mess)
	if err != nil {
		log.Fatal(err)
	}
	api := slack.New(mess.BotToken)
	wg := &sync.WaitGroup{}
	AllChannels := make(map[string]*sync.Mutex)
	for _, channel := range mess.Channels {
		if _, ok := AllChannels[channel.ChanName]; !ok {
			AllChannels[channel.ChanName] = &sync.Mutex{}
		}
		wg.Add(1)
		go sendToChan(wg, AllChannels[channel.ChanName], channel.ChanName, channel.Text, api)
	}
	wg.Wait()
}