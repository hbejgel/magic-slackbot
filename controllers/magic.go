package controllers

import (
	"github.com/astaxie/beego"
	"github.com/hbejgel/magic-slackbot/models/magic"
	"github.com/hbejgel/magic-slackbot/models/slack"
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
	query := this.Input().Get("text")
	cards, err := magic.GetCardsWithText(query)
	if err != nil {
		this.Data["json"] = err
	} else {
		image, err := cards.GetRandomCardImage()
		println("Image:", image)
		if err != nil {
			this.Data["json"] = err
		} else {
			slack.SendImage(image, this.Input())
		}
	}
}
