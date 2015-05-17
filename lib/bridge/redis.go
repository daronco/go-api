package bridge

import (
	"encoding/json"
	"fmt"
	"github.com/bigbluebutton/go-api/models"
	"github.com/garyburd/redigo/redis"
)

type RedisBridge struct {
	redisPool      *redis.Pool
	redisAddress   string
	maxConnections int
}

func (this *RedisBridge) Start() {
	this.redisPool = redis.NewPool(func() (redis.Conn, error) {
		c, err := redis.Dial("tcp", this.redisAddress)
		if err != nil {
			return nil, err
		}
		return c, err
	}, this.maxConnections)
	fmt.Printf("--- redis connected\n")
	return
}

func (this *RedisBridge) Stop() {
	fmt.Printf("--- stopping redis\n")
	this.redisPool.Close()
}

func (this *RedisBridge) GetMeeting(meetingId string) (models.Meeting, error) {
	var c = this.redisPool.Get()
	var meeting models.Meeting

	response, err := redis.Bytes(c.Do("HGET", meetingId, "value"))
	fmt.Printf("--- redis got: %s\n", response)
	if err != nil {
		return meeting, err
	} else {
		err = json.Unmarshal(response, &meeting)
		if err != nil {
			fmt.Printf("--- not a Meeting object\n")
			return meeting, err
		} else {
			fmt.Printf("--- meeting loaded successfully\n")
			return meeting, nil
		}
	}
}
