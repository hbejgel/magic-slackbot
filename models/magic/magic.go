package magic

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
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

func GetCardsWithText(text string) (CardsResponse, error) {
	cards, err := GeneralCardGetter("name=" + text)
	if err != nil {
		return CardsResponse{}, err
	}
	if len(cards.Cards) > 0 {
		return cards, nil
	}
	arguments := strings.Split(text, " ")
	if len(arguments) == 1 {
		query := "colors=[\"" + arguments[0] + "\"]"
		cards, err := GeneralCardGetter(query)
		if err != nil {
			return CardsResponse{}, err
		}
		random_page := rand.Intn((cards.Total / cards.PerPage))
		return GeneralCardGetter(query + fmt.Sprintf("&page=%v", random_page))
	} else if len(arguments) == 2 {
		query := "colors=[\"" + arguments[0] + "\"]&cmc=" + arguments[1]
		cards, err := GeneralCardGetter(query)
		if err != nil {
			return CardsResponse{}, err
		}
		random_page := rand.Intn((cards.Total / cards.PerPage))
		return GeneralCardGetter(query + fmt.Sprintf("&page=%v", random_page))
	}
	return CardsResponse{}, errors.New("Invalid Arguments")
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

func (this CardsResponse) GetRandomCardImage() (string, error) {
	total_cards := len(this.Cards)
	if total_cards == 0 {
		return "", errors.New("Not found")
	}
	image_link, ok := this.Cards[rand.Intn(total_cards)].Images["gatherer"]
	if ok {
		return image_link, nil
	}
	return "", errors.New("No Image")
}
