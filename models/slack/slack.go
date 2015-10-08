package slack

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
)

type WebHook struct {
	Channel     string       `json:"channel,omitempty"`
	Attachments []Attachment `json:"attachments"`
}

type Attachment struct {
	Fallback   string `json:"fallback"`
	Title      string `json:"title"`
	Image      string `json:"image_url"`
	AuthorName string `json:"author_name,omitempty"`
}

const (
	zuerinhaUrl   = "https://hooks.slack.com/services/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"
	recurrenceUrl = "https://hooks.slack.com/services/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"
)

func SendImage(image string, info url.Values) {
	channel := info.Get("channel_id")
	if channel == "" {
		return
	}
	var wh WebHook
	wh.Channel = channel
	attachment := Attachment{}
	attachment.Fallback = "Magic Images!"
	attachment.Image = image
	attachment.Title = "Magic Bot!"
	attachment.AuthorName = info.Get("user_name")
	wh.Attachments = append(wh.Attachments, attachment)
	println("Channel destination:", channel)
	data, _ := json.Marshal(wh)
	http.Post(zuerinhaUrl, "application/json", bytes.NewBuffer(data))
}
