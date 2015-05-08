package bridge

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/bigbluebutton/go-api/lib/bridge/redis"
)

var (
	redisBridge redis.RedisBridge
)

func init() {
	redisBridge = redis.RedisBridge{}
}

func Start() {
	fmt.Printf("--- init bridge\n")
	var port, _ = beego.AppConfig.Int("RedisPort")
	var redisAddress = fmt.Sprintf("%s:%d", beego.AppConfig.String("RedisAddress"), port)
	var maxConnections, _ = beego.AppConfig.Int("RedisMaxConnections")

	redisBridge.Start(redisAddress, maxConnections)
	fmt.Printf("--- bridge init done\n")
}

func Stop() {
	redisBridge.Stop()
}

func Save(key string, obj interface{}) (bool, error) {
	marshalled, _ := json.Marshal(obj)
	fmt.Printf("--- bridge saving: %s\n", marshalled)
	return redisBridge.Send(key, marshalled)
}

func Get(key string) ([]byte, error) {
	var response, err = redisBridge.Get(key)
	fmt.Printf("--- bridge read: %v\n", response)
	if err != nil {
		return nil, err
	} else {
		return response, nil
	}
}
