package main

import (
	_ "github.com/bigbluebutton/api-labs/docs"
	_ "github.com/bigbluebutton/api-labs/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
