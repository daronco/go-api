package main

import (
	_ "github.com/bigbluebutton/go-api/docs"
	_ "github.com/bigbluebutton/go-api/routers"

	"fmt"
	"github.com/astaxie/beego"
	"github.com/bigbluebutton/go-api/lib/bridge"
	// "github.com/bigbluebutton/go-api/models"
)

func main() {
	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"
	}

	bridge.Start()
	defer bridge.Stop()

	// meeting := models.Meeting{"meeting1", "My Meeting", ""}
	// result, err := meeting.Save()
	// fmt.Printf("result: %v\n", result)
	// fmt.Printf("error: %v\n", err)
	// fmt.Printf("--- addMeeting ok")

	beego.Run()
	fmt.Printf("--- run ok")
}
