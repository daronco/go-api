package redis

import (
	// "encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

type RedisBridge struct {
	redisPool *redis.Pool
}

func (this *RedisBridge) Start(redisAddress string, maxConnections int) {
	this.redisPool = redis.NewPool(func() (redis.Conn, error) {
		c, err := redis.Dial("tcp", redisAddress)
		if err != nil {
			return nil, err
		}
		return c, err
	}, maxConnections)
	fmt.Printf("--- redis connected\n")
	return
}

func (this *RedisBridge) Stop() {
	fmt.Printf("--- stopping redis\n")
	this.redisPool.Close()
}

func (this *RedisBridge) Send(key string, value []byte) (bool, error) {
	var c = this.redisPool.Get()

	fmt.Printf("--- redis sending: %s\n", value)

	_, err := c.Do("HSET", key, "value", value)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

func (this *RedisBridge) Get(key string) ([]byte, error) {
	var c = this.redisPool.Get()

	reply, err := redis.Bytes(c.Do("HGET", key, "value"))
	fmt.Printf("--- redis got: %s\n", reply)
	if err != nil {
		return nil, err
	} else {
		return reply, nil
	}
}
