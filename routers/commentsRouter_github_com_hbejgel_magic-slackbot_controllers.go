package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/hbejgel/magic-slackbot/controllers:MagicController"] = append(beego.GlobalControllerRouter["github.com/hbejgel/magic-slackbot/controllers:MagicController"],
		beego.ControllerComments{
			"Get",
			`/magic`,
			[]string{"get"},
			nil})

}
