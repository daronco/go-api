package main

import (
	_ "github.com/bigbluebutton/go-api/docs"
	_ "github.com/bigbluebutton/go-api/routers"

	"github.com/astaxie/beego"
	"github.com/bigbluebutton/go-api/lib/redispubsub"
)

func main() {
	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"

	}

	go redispubsub.Connect()

	beego.Run()
}
