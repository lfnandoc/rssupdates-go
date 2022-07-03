package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type DiscordMessage struct {
	Username   string         `json:"username,omitempty"`
	AvatarURL  string         `json:"avatar_url,omitempty"`
	Content    string         `json:"content,omitempty"`
	Embeds     []Embed        `json:"embeds,omitempty"`
	Components []ComponentRow `json:"components,omitempty"`
}

type Embed struct {
	Title       string    `json:"title"`
	Color       int       `json:"color"`
	Description string    `json:"description"`
	Timestamp   time.Time `json:"timestamp"`
	URL         string    `json:"url"`
	Author      struct {
	} `json:"author,omitempty"`
	Image     Image     `json:"image"`
	Thumbnail Thumbnail `json:"thumbnail"`
	Footer    struct {
	} `json:"footer,omitempty"`
	Fields []Field `json:"fields"`
}

type Image struct {
	URL string `json:"url"`
}

type Thumbnail struct {
	URL string `json:"url"`
}

type Field struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}

type ComponentRow struct {
	Type       int         `json:"type"`
	Components []Component `json:"components"`
}

type Component struct {
	Type  int    `json:"type"`
	Style int    `json:"style"`
	Label string `json:"label"`
	URL   string `json:"url"`
}

func SendDiscordMessageToWebhook(discordMessage DiscordMessage) (err error) {

	data, err := json.Marshal(discordMessage)
	if err != nil {
		return err
	}

	response, err := http.Post(Configs.DiscordWebhook, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	if response.StatusCode != 204 {
		return fmt.Errorf("discord webhook returned %d", response.StatusCode)
	}

	return nil
}
