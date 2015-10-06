package controllers

import (
	//"encoding/json"
	"github.com/astaxie/beego"
	"github.com/hbejgel/magic-slackbot/models"
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
		this.Ctx.ResponseWriter.Write([]byte(cards.GetRandomCardImage()))
		//this.Data["json"] = cards.GetRandomCardImage()
	}
}
