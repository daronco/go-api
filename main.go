package main

import (
	_ "github.com/bigbluebutton/api-labs/docs"
	_ "github.com/bigbluebutton/api-labs/routers"

	"github.com/astaxie/beego"
	"github.com/bigbluebutton/api-labs/lib/redispubsub"
)

func main() {
	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"

	}

	go redispubsub.Connect()

	beego.Run()
}
