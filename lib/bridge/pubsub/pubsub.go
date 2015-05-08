package pubsub

import (
	// "encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

type PubSubBridge struct {
	redisPool *redis.Pool
	psc       redis.PubSubConn
}

func (this *PubSubBridge) Start(redisAddress string, maxConnections int) {
	this.redisPool = redis.NewPool(func() (redis.Conn, error) {
		c, err := redis.Dial("tcp", redisAddress)
		if err != nil {
			return nil, err
		}
		return c, err
	}, maxConnections)
	fmt.Printf("--- redis connected\n")

	this.psc = redis.PubSubConn{Conn: this.redisPool.Get()}
	this.psc.Subscribe("bigbluebutton:from-bbb-apps:meeting")

	return
}

// func (this *PubSubBridge)
// 	psc := redis.PubSubConn{Conn: this.redisPool.Get()}
// 	psc.Subscribe("bigbluebutton:from-bbb-apps:meeting")

// 	for {
// 		switch event := psc.Receive().(type) {
// 		case redis.Message:
// 			fmt.Printf("[redis message] %s: message: %s\n", event.Channel, event.Data)
// 			this.treatMessage([]byte(event.Data))
// 		case redis.Subscription:
// 			fmt.Printf("[redis subscription] %s: %s %d\n", event.Channel, event.Kind, event.Count)
// 		case error:
// 			fmt.Printf("error: %v\n", event)
// 			return
// 		}
// 	}
// 	// return
// }

func (this *PubSubBridge) Stop() {
	fmt.Printf("--- stopping redis\n")
	this.redisPool.Close()
}

func (this *PubSubBridge) Get(key string) ([]byte, error) {

	// TODO: send a message to bbb-apps asking for the data

	for {
		switch event := this.psc.Receive().(type) {
		case redis.Message:
			fmt.Printf("[redis message] %s: message: %s\n", event.Channel, event.Data)
			return []byte(event.Data), nil
		case redis.Subscription:
			fmt.Printf("[redis subscription] %s: %s %d\n", event.Channel, event.Kind, event.Count)
		case error:
			return nil, event
		}
	}
}

// func (this *PubSubBridge) treatMessage(msg []byte) {
// 	var deserialized = &messages.MeetingCreatedMessage{}
// 	var err = json.Unmarshal(msg, &deserialized)

// 	if err == nil {
// 		fmt.Println("--- deserialized data: ", deserialized)
// 	}
// }
