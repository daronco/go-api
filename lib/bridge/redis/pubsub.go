package redis

// psc := redis.PubSubConn{Conn: redisPool.Get()}
// psc.Subscribe("bigbluebutton:from-bbb-apps:meeting")

// for {
// 	switch event := psc.Receive().(type) {
// 	case redis.Message:
// 		fmt.Printf("[redis message] %s: message: %s\n", event.Channel, event.Data)
// 		treatMessage([]byte(event.Data))
// 	case redis.Subscription:
// 		fmt.Printf("[redis subscription] %s: %s %d\n", event.Channel, event.Kind, event.Count)
// 	case error:
// 		fmt.Printf("error: %v\n", event)
// 		return
// 	}
// }

// func treatMessage(msg []byte) {
// 	var deserialized = &MeetingCreatedMessage{}
// 	var err = json.Unmarshal(msg, &deserialized)

// 	if err == nil {
// 		fmt.Println("deserialized data: ", deserialized)
// 	}
// }
