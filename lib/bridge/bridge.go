package bridge

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/bigbluebutton/go-api/models"
)

var (
	bridgeObj           BridgeInterface
	bridgeType          string
	redisAddress        string
	redisMaxConnections int
)

func init() {
	bridgeType = beego.AppConfig.String("bridgetype")

	port, _ := beego.AppConfig.Int("RedisPort")
	redisAddress = fmt.Sprintf("%s:%d", beego.AppConfig.String("RedisAddress"), port)
	redisMaxConnections, _ = beego.AppConfig.Int("RedisMaxConnections")
}

func Start() {
	fmt.Printf("--- init bridge: %s\n", bridgeType)

	switch bridgeType {
	case "pubsub":
		fmt.Printf("--- starting pubsub bridge\n")
		bridgeObj = &PubSubBridge{redisAddress: redisAddress, maxConnections: redisMaxConnections}
		bridgeObj.Start()
	case "redis":
		fmt.Printf("--- starting redis bridge\n")
		bridgeObj = &RedisBridge{redisAddress: redisAddress, maxConnections: redisMaxConnections}
		bridgeObj.Start()
	default:
		panic("unrecognized bridge type")
	}
	fmt.Printf("--- bridge start done\n")
}

func Stop() {
	bridgeObj.Stop()
	fmt.Printf("--- bridge stop done\n")
}

func GetMeeting(meetingId string) (models.Meeting, error) {
	meeting, err := bridgeObj.GetMeeting(meetingId)
	fmt.Printf("--- bridge read: %v\n", meeting)
	if err != nil {
		return meeting, err
	} else {
		return meeting, nil
	}
}
