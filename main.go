package main

import (
	_ "github.com/hbejgel/magic-slackbot/docs"
	_ "github.com/hbejgel/magic-slackbot/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
