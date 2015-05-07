package redispubsub

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var (
	redisPool      *redis.Pool
	redisAddress   = flag.String("redis-address", ":6379", "Address to the Redis server")
	maxConnections = flag.Int("max-connections", 10, "Max connections to Redis")
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

func Connect() {
	flag.Parse()

	redisPool := redis.NewPool(func() (redis.Conn, error) {
		c, err := redis.Dial("tcp", *redisAddress)

		if err != nil {
			return nil, err
		}

		return c, err
	}, *maxConnections)

	psc := redis.PubSubConn{Conn: redisPool.Get()}

	psc.Subscribe("bigbluebutton:from-bbb-apps:meeting")

	for {
		switch event := psc.Receive().(type) {
		case redis.Message:
			fmt.Printf("[redis message] %s: message: %s\n", event.Channel, event.Data)
			treatMessage([]byte(event.Data))
		case redis.Subscription:
			fmt.Printf("[redis subscription] %s: %s %d\n", event.Channel, event.Kind, event.Count)
		case error:
			fmt.Printf("error: %v\n", event)
			return
		}
	}
}

func treatMessage(msg []byte) {
	var deserialized = &MeetingCreatedMessage{}
	var err = json.Unmarshal(msg, &deserialized)

	if err == nil {
		fmt.Println("deserialized data: ", deserialized)
	}
}
