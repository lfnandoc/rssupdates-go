package main

import (
	"encoding/json"
	"io/ioutil"
)

var Configs Configuration

type Configuration struct {
	DiscordWebhook string
	RssFeed        string
}

func SetupConfiguration() {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	var config Configuration
	err = json.Unmarshal(file, &config)
	if err != nil {
		panic(err)
	}

	Configs = config
}
