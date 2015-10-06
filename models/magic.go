package models

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
)

const (
	baseApiUrl = "http://api.mtgapi.com/v2/cards"
)

type CardsResponse struct {
	Queries []Query           `json:"query"`
	Cards   []Card            `json:"cards"`
	Total   int               `json:"total"`
	PerPage int               `json:"perPage"`
	Links   map[string]string `json:"links"`
}

type Query struct {
	Command     string `json:"command"`
	Key         string `json:"key"`
	Conditional string `json:"conditional"`
	Value       string `json:"value"`
}

type Card struct {
	Artist string            `json:"artist"`
	Cmc    int               `json:"cmc"`
	Images map[string]string `json:"images"`
	Name   string            `json:"name"`
}

type WebHook struct {
	Channel     string       `json:"channel,omitempty"`
	Attachments []Attachment `json:"attachments"`
}

type Attachment struct {
	Fallback string `json:"fallback"`
	Title    string `json:"title"`
	Image    string `json:"image_url"`
}

func GeneralCardGetter(query string) (CardsResponse, error) {
	endpoint, err := url.Parse(baseApiUrl)
	endpoint.RawQuery = query
	println("Outgoing reqeust: ", endpoint.String())
	resp, err := http.Get(endpoint.String())
	if err != nil {
		return CardsResponse{}, err
	}

	if resp.StatusCode >= 400 {
		return CardsResponse{}, errors.New("Response Error")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return CardsResponse{}, err
	}

	var response CardsResponse
	json.Unmarshal(body, &response)

	return response, nil
}

func (this CardsResponse) GetRandomCardImage() (WebHook, error) {
	total_cards := len(this.Cards)
	if total_cards == 0 {
		return WebHook{}, errors.New("Not found")
	}
	image_link, ok := this.Cards[rand.Intn(total_cards)].Images["gatherer"]
	if ok {
		var wh WebHook
		attachment := Attachment{}
		attachment.Fallback = "Magic Images!"
		attachment.Image = image_link
		attachment.Title = "Magic Bot!"
		wh.Attachments = append(wh.Attachments, attachment)
		return wh, nil
	}
	return WebHook{}, errors.New("No Image")
}
