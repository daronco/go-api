package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	
	beego.GlobalControllerRouter["github.com/bigbluebutton/api-labs/controllers:MeetingController"] = append(beego.GlobalControllerRouter["github.com/bigbluebutton/api-labs/controllers:MeetingController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/bigbluebutton/api-labs/controllers:MeetingController"] = append(beego.GlobalControllerRouter["github.com/bigbluebutton/api-labs/controllers:MeetingController"],
		beego.ControllerComments{
			"Get",
			`/:meetingId`,
			[]string{"get"},
			nil})

}
