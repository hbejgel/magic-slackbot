package controllers

import (
	//"encoding/json"
	"bytes"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/hbejgel/magic-slackbot/models"
	"net/http"
)

// Operations about magic cards
type MagicController struct {
	beego.Controller
}

// @Title get card
// @Description get card
// @Success 200 card
// @Failure 404 card not found
// @router /magic [get]
func (this *MagicController) Get() {
	card_name := this.Input().Get("text")
	cards, err := models.GeneralCardGetter("name=" + card_name)
	if err != nil {
		this.Data["json"] = err
	} else {
		wh, err := cards.GetRandomCardImage()
		if err != nil {
			this.Data["json"] = err
		} else {
			channel := this.Input().Get("channel_id")
			if channel != "" {
				wh.Channel = channel
			}
			println("Channel destination:", channel)
			data, _ := json.Marshal(wh)
			http.Post("https://hooks.slack.com/services/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", "application/json", bytes.NewBuffer(data))
			println("Done")
		}
	}
}
