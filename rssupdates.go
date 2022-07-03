package main

import (
	"encoding/xml"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/robfig/cron"
)

func main() {
	SetupDatabase()
	SetupConfiguration()

	c := cron.New()
	c.AddFunc("@every 15m", func() {
		CheckForNewPosts()
	})
	c.Start()

	select {}
}

func CheckForNewPosts() {
	response, err := http.Get(Configs.RssFeed)
	if err != nil {
		panic(err)
	}

	if response.StatusCode != 200 {
		panic(response.StatusCode)
	}

	var rss Rss
	err = xml.NewDecoder(response.Body).Decode(&rss)
	if err != nil {
		panic(err)
	}

	for _, item := range rss.Channel.Item {
		if CheckIfNewPost(item.Guid.Text) {

			post := Post{
				Guid: item.Guid.Text,
			}
			DB.Create(&post)

			postTime, _ := time.Parse(time.RFC1123Z, item.PubDate)

			embed := Embed{
				Title:       item.Title,
				URL:         item.Link,
				Description: item.Description,
				Thumbnail:   Thumbnail{URL: GetFirstImgOfHtml(item.Encoded)},
				Timestamp:   postTime,
			}

			discordMessage := DiscordMessage{
				Embeds: []Embed{embed},
			}

			SendDiscordMessageToWebhook(discordMessage)
		}
	}

}

func CheckIfNewPost(guid string) bool {
	var post Post
	DB.Find(&post, "guid = ?", guid)
	return post.ID == 0
}

func GetFirstImgOfHtml(htmlstr string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlstr))
	if err != nil {
		panic(err)
	}

	img := doc.Find("img")
	if img.Length() == 0 {
		return ""
	}

	return img.First().AttrOr("src", "")
}
