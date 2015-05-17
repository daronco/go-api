package messages

import (
	"encoding/json"
	"github.com/bigbluebutton/go-api/models"
	// "fmt"
)

type MessageHeader struct {
	Timestamp uint64 `json:"timestamp"`
	Name      string `json:"name"`
	Version   string `json:"version"`
}

type MeetingCreatedPayload struct {
	Duration                int    `json:"duration"`
	CreateDate              string `json:"create_data"`
	Name                    string `json:"name"`
	CreateTime              uint64 `json:"create_time"`
	ModeratorPass           string `json:"moderator_pass"`
	AllowStartStopRecording bool   `json:"allow_start_stop_recording"`
	VoiceConf               string `json:"voice_conf"`
	Recorded                bool   `json:"recorded"`
	ExternalMeetingId       string `json:"external_meeting_id"`
	MeetingId               string `json:"meeting_id"`
	ViewerPass              string `json:"viewer_pass"`
	AutoStartRecording      bool   `json:"allow_start_recording"`
}

type MeetingCreatedMessage struct {
	Header  MessageHeader         `json:"header"`
	Payload MeetingCreatedPayload `json:"payload"`
}

func Parse(event []byte) (interface{}, error) {

	// TODO: find the type of message by reading event's header and use the correct struct

	deserialized := &MeetingCreatedMessage{}
	err := json.Unmarshal(event, &deserialized)

	if err != nil {
		return nil, err
	} else {
		return deserialized, nil
	}
}

func (this *MeetingCreatedMessage) ToMeeting() (models.Meeting) {
	// TODO: temporary
	return models.Meeting{"meeting1", "My Meeting", ""}
}
