package bridge

import (
	"github.com/bigbluebutton/go-api/models"
)

type BridgeInterface interface {
	GetMeeting(meetingId string) (models.Meeting, error)
	Start()
	Stop()
}
