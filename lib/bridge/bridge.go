package bridge

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/bigbluebutton/go-api/lib/bridge/pubsub"
	"github.com/bigbluebutton/go-api/lib/bridge/redis"
	"github.com/bigbluebutton/go-api/lib/messages"
)

var (
	redisBridge  redis.RedisBridge
	pubsubBridge pubsub.PubSubBridge
	bridgeType   string

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
		pubsub_start()
	case "redis":
		redis_start()
	default:
		panic("unrecognized bridge type")
	}
	fmt.Printf("--- bridge start done\n")
}

func Stop() {
	switch bridgeType {
	case "pubsub":
		pubsub_stop()
	case "redis":
		redis_stop()
	default:
		panic("unrecognized bridge type")
	}
	fmt.Printf("--- bridge stop done\n")
}

func Save(key string, obj interface{}) (bool, error) {
	switch bridgeType {
	case "pubsub":
		return pubsub_save(key, obj)
	case "redis":
		return redis_save(key, obj)
	default:
		panic("unrecognized bridge type")
	}
}

func Get(key string) ([]byte, error) {
	switch bridgeType {
	case "pubsub":
		return pubsub_get(key)
	case "redis":
		return redis_get(key)
	default:
		panic("unrecognized bridge type")
	}
}

func pubsub_start() {
	fmt.Printf("--- starting pubsub bridge\n")
	pubsubBridge = pubsub.PubSubBridge{}
	pubsubBridge.Start(redisAddress, redisMaxConnections)
}

func redis_start() {
	fmt.Printf("--- starting redis bridge\n")
	redisBridge = redis.RedisBridge{}
	redisBridge.Start(redisAddress, redisMaxConnections)
}

func pubsub_stop() {
	pubsubBridge.Stop()
}

func redis_stop() {
	redisBridge.Stop()
}

func redis_save(key string, obj interface{}) (bool, error) {
	marshalled, _ := json.Marshal(obj)
	fmt.Printf("--- bridge saving: %s\n", marshalled)
	return redisBridge.Send(key, marshalled)
}

func pubsub_save(key string, obj interface{}) (bool, error) {
	// TODO
	return true, nil
}

func redis_get(key string) ([]byte, error) {
	var response, err = redisBridge.Get(key)
	fmt.Printf("--- bridge read: %v\n", response)
	if err != nil {
		return nil, err
	} else {
		return response, nil
	}
}

func pubsub_get(key string) ([]byte, error) {
	response, err := pubsubBridge.Get(key)
	fmt.Printf("--- bridge read: %v\n", response)

	// TOOD: response here is an event from redis, has to be converted to a model

	if err != nil {
		return nil, err
	} else {
		var parsed interface{}
		parsed, err = messages.Parse(response)
		if parsed != nil {
			return getBytes(parsed)
		} else {
			return nil, err
		}
	}
}

func getBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
