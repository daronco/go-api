package controllers

import (
	"github.com/bigbluebutton/go-api/lib/bridge"
	"github.com/astaxie/beego"
)

// Operations over meetings
type MeetingController struct {
	beego.Controller
}

// @Title GetAll
// @Description get all meetings
// @Success 200 {object} models.Meeting
// @router / [get]
func (this *MeetingController) GetAll() {
	this.Data["json"] = nil //bridge.GetAllMeetings()
	this.ServeJson()
}

// @Title Get
// @Description get meeting by meetingId
// @Param	meetingId		path 	string	true		"The key to find the meeting"
// @Success 200 {object} models.Meeting
// @Failure 403 :meetingId is empty
// @router /:meetingId [get]
func (this *MeetingController) Get() {
	meetingId := this.GetString(":meetingId")
	if meetingId != "" {
		meeting, err := bridge.GetMeeting(meetingId)
		if err != nil {
			this.Data["json"] = err
		} else {
			this.Data["json"] = meeting
		}
	}
	this.ServeJson()
}
