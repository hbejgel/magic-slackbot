package main

import (
	"github.com/astaxie/beego"
	_ "github.com/hbejgel/magic-slackbot/routers"
	"os"
	"strconv"
)

func main() {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err == nil {
		beego.HttpPort = port
	}
	beego.Run()
}
