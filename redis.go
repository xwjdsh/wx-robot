package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

var redisPool *redis.Pool

func initRedis() {
	redisPool = &redis.Pool{
		MaxIdle:   80,
		MaxActive: 1200,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", *redisFlag)
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}
}

func redisError(err error) string {
	return fmt.Sprintf("Redis error:%s", err.Error())
}
