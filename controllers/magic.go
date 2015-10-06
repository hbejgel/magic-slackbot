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
func (this *MagicController) Post() {
	query := this.Ctx.Request.URL.RawQuery
	cards, err := models.GeneralCardGetter(query)
	if err != nil {
		this.Data["json"] = err
	} else {
		this.Data["json"] = cards
	}
	this.ServeJson()
}
