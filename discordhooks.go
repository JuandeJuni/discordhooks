package discordhooks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Field struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}
type Thumbnail struct {
	Url string `json:"url"`
}
type Footer struct {
	Text     string `json:"text"`
	Icon_url string `json:"icon_url"`
}
type Embed struct {
	Title       string    `json:"title"`
	Url         string    `json:"url"`
	Description string    `json:"description"`
	Color       int       `json:"color"`
	Thumbnail   Thumbnail `json:"thumbnail"`
	Footer      Footer    `json:"footer"`
	Fields      []Field   `json:"fields"`
}
type Attachment struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	Filename    string `json:"filename"`
}
type Hook struct {
	Username    string       `json:"username"`
	Avatar_url  string       `json:"avatar_url"`
	Content     string       `json:"content"`
	Embeds      []Embed      `json:"embeds"`
	Attachments []Attachment `json:"attachments"`
}

func ExecuteWebhook(link string, data []byte) {

	req, err := http.NewRequest("POST", link, bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == 429 {
		fmt.Println("Rate limit reached")
		time.Sleep(time.Second * 5)
		ExecuteWebhook(link, data)
	}
}
func SendEmbeds(link string, embeds []Embed) {
	hook := Hook{
		Embeds: embeds,
	}
	payload, err := json.Marshal(hook)
	if err != nil {
		log.Fatal(err)
	}
	executeWebhook(link, payload)

}
func SendEmbed(link string, embeds Embed) {
	hook := Hook{
		Embeds: []Embed{embeds},
	}
	payload, err := json.Marshal(hook)
	if err != nil {
		log.Fatal(err)
	}
	executeWebhook(link, payload)

}
