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
		attachment, err := cards.GetRandomCardImage()
		if err != nil {
			this.Data["json"] = err
		} else {
			println(attachment.Att[0].Image)
			data, _ := json.Marshal(attachment)
			http.Post("https://hooks.slack.com/services/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", "application/json", bytes.NewBuffer(data))
			println("Done")
		}
	}
}
