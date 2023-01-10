package discordhooks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type FieldS struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline,omitempty"`
}
type ThumbnailS struct {
	Url string `json:"url"`
}
type FooterS struct {
	Text    string `json:"text"`
	IconURL string `json:"icon_url"`
}
type AuthorS struct {
	Name string `json:"name"`
}

type EmbedS struct {
	Title     string     `json:"title"`
	URL       string     `json:"url"`
	Color     int        `json:"color"`
	Fields    []FieldS   `json:"fields"`
	Author    AuthorS    `json:"author"`
	Footer    FooterS    `json:"footer"`
	Timestamp time.Time  `json:"timestamp"`
	Thumbnail ThumbnailS `json:"thumbnail"`
}
type Attachment struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Filename    string `json:"filename"`
}
type NewHook struct {
	Content     interface{}  `json:"content"`
	Embeds      []EmbedS     `json:"embeds"`
	Username    string       `json:"username"`
	AvatarURL   string       `json:"avatar_url"`
	Attachments []Attachment `json:"attachments"`
}

func ExecuteWebhook(link string, data []byte) error {

	req, err := http.NewRequest("POST", link+"?wait=true", bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authority", "discord.com")
	req.Header.Set("accept", "application/json")
	req.Header.Set("accept-language", "en")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("pragma", "no-cache")
	req.Header.Set("sec-ch-ua", `"Not?A_Brand";v="8", "Chromium";v="108", "Google Chrome";v="108"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "cross-site")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		fmt.Printf("%s\n", bodyText)
	}
	if resp.StatusCode == 429 {
		fmt.Println("Rate limit reached")
		time.Sleep(time.Second * 5)
		ExecuteWebhook(link, data)
	}
	return err
}

func SendEmbeds(link string, embeds []EmbedS) error {
	hook := NewHook{
		Embeds: embeds,
	}

	payload, err := json.Marshal(hook)
	if err != nil {
		log.Fatal(err)
	}
	err = ExecuteWebhook(link, payload)
	return err

}

// func SendEmbedWithFile(link string, embeds Embed, filepath string) error {

// 	hook := Hook{
// 		Embeds:      []Embed{embeds},
// 		Attachments: []Attachment{{ID: "0", Description: "thumbnail", Filename: "myfilename.png"}},
// 		Content:     "test",
// 	}
// 	bytepayload, err := json.Marshal(hook)
// 	if err != nil {
// 		return err
// 	}
// 	fileDir, _ := os.Getwd()
// 	fileName := "myfilename.png"
// 	filePath := path.Join(fileDir, fileName)

// 	file, _ := os.Open(filePath)
// 	defer file.Close()

// 	body := &bytes.Buffer{}
// 	writer := multipart.NewWriter(body)
// 	metadataHeader := textproto.MIMEHeader{}
// 	metadataHeader.Set("Content-Type", "application/json; charset=UTF-8")
// 	part1, _ := writer.CreatePart(metadataHeader)
// 	part1.Write(bytepayload)
// 	part, _ := writer.CreateFormFile("file", "myfilename.png")
// 	io.Copy(part, file)
// 	writer.Close()

// 	r, _ := http.NewRequest("POST", link, body)
// 	r.Header.Add("Content-Type", writer.FormDataContentType())
// 	client := &http.Client{}
// 	resp, err := client.Do(r)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer resp.Body.Close()
// 	bodyText, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if resp.StatusCode != 200 {
// 		fmt.Printf("%s\n", bodyText)
// 	}
// 	if resp.StatusCode == 429 {
// 		fmt.Println("Rate limit reached")
// 		time.Sleep(time.Second * 5)
// 		SendEmbedWithFile(link, embeds, filepath)
// 	}
// 	return err
// 	// body := &bytes.Buffer{}
// 	// // Creates a new multipart Writer with a random boundary
// 	// // writing to the empty buffer
// 	// writer := multipart.NewWriter(body)
// 	// metadataHeader := textproto.MIMEHeader{}
// 	// metadataHeader.Set("Content-Type", "application/json; charset=UTF-8")
// 	// part, err := writer.CreatePart(metadataHeader)
// 	// if err != nil {
// 	// 	return err
// 	// }
// 	// // Write the part body
// 	// part.Write(bytepayload)
// 	// mediaData, err := ioutil.ReadFile("image.png")
// 	// if err != nil {
// 	// 	return err
// 	// }
// 	// mediaHeader := textproto.MIMEHeader{}
// 	// mediaHeader.Set("Content-Type", "image/png")

// 	// mediaPart, err := writer.CreatePart(mediaHeader)
// 	// if err != nil {
// 	// 	return err
// 	// }
// 	// io.Copy(mediaPart, bytes.NewReader(mediaData))

// 	// // Finish constructing the multipart request body
// 	// writer.Close()

// 	// // dat, err := os.ReadFile(filepath)
// 	// // if err != nil {
// 	// // 	log.Fatal(err)
// 	// // }
// 	// // payload := `--boundary` + "\n" + `Content-Disposition: form-data; name="payload_json"` + "\n" + `Content-Type: application/json` + "\n" + `` + "\n" + ``

// 	// // payload += string(bytepayload)
// 	// // payload += `` + "\n" + `--boundary` + "\n" + `Content-Disposition: form-data; name="files[0]"; filename="myfilename.png"` + "\n" + `Content-Type: image/png` + "\n" + `` + "\n" + `` + string(dat) + `` + "\n" + `--boundary`
// 	// // fmt.Println(payload)
// 	// // if err != nil {
// 	// // 	log.Fatal(err)
// 	// // }
// 	// // err = ExecuteWebhookFile(link, body,writer)
// 	// req, err := http.NewRequest("POST", link, bytes.NewReader(body.Bytes()))
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// contentType := fmt.Sprintf("multipart/form-data; boundary=%s", writer.Boundary())
// 	// req.Header.Set("Content-Type", contentType)
// 	// // Content-Length must be the total number of bytes in the request body.
// 	// req.Header.Set("Content-Length", fmt.Sprintf("%d", body.Len()))
// 	// client := &http.Client{}
// 	// resp, err := client.Do(req)
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// defer resp.Body.Close()
// 	// bodyText, err := ioutil.ReadAll(resp.Body)
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// if resp.StatusCode != 200 {
// 	// 	fmt.Printf("%s\n", bodyText)
// 	// }
// 	// if resp.StatusCode == 429 {
// 	// 	fmt.Println("Rate limit reached")
// 	// 	time.Sleep(time.Second * 5)
// 	// 	SendEmbedWithFile(link, embeds, filepath)
// 	// }
// 	// return err

// }
func SendEmbed(username string, sitelogo string, link string, embeds EmbedS) error {
	hook := NewHook{
		Username:  username,
		AvatarURL: sitelogo,
		Embeds:    []EmbedS{embeds},
	}
	payload, err := json.Marshal(hook)
	if err != nil {
		log.Fatal(err)
	}
	err = ExecuteWebhook(link, payload)
	return err

}
