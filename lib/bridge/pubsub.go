package bridge

import (
	"fmt"
	"github.com/bigbluebutton/go-api/models"
	"github.com/bigbluebutton/go-api/lib/messages"
	"github.com/garyburd/redigo/redis"
)

type PubSubBridge struct {
	redisPool      *redis.Pool
	psc            redis.PubSubConn
	redisAddress   string
	maxConnections int
}

func (this *PubSubBridge) Start() {
	this.redisPool = redis.NewPool(func() (redis.Conn, error) {
		c, err := redis.Dial("tcp", this.redisAddress)
		if err != nil {
			return nil, err
		}
		return c, err
	}, this.maxConnections)
	fmt.Printf("--- redis connected\n")

	this.psc = redis.PubSubConn{Conn: this.redisPool.Get()}
	this.psc.Subscribe("bigbluebutton:from-bbb-apps:meeting")

	return
}

func (this *PubSubBridge) Stop() {
	fmt.Printf("--- stopping redis\n")
	this.redisPool.Close()
}

func (this *PubSubBridge) GetMeeting(meetingId string) (models.Meeting, error) {
	var meeting models.Meeting

	// TODO: send a message to bbb-apps asking for the data

	for {
		switch event := this.psc.Receive().(type) {
		case redis.Message:
			fmt.Printf("[redis message] %s: message: %s\n", event.Channel, event.Data)

			// parse the message received into a struct
			message, err := messages.Parse(event.Data)
			if err != nil {
				return meeting, err
			} else {
				// TODO: check 'message' is the correct message

				// get the meeting from the message
				if m, ok := message.(messages.MeetingCreatedMessage); ok {
					return m.ToMeeting(), nil
				} else {
					return meeting, nil
				}
			}

		case redis.Subscription:
			fmt.Printf("[redis subscription] %s: %s %d\n", event.Channel, event.Kind, event.Count)
		case error:
			return meeting, event
		}
	}
}
