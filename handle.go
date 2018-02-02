package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/garyburd/redigo/redis"
)

type (
	handler interface {
		Do(*WxMessage) error
		Desc() string
		Key() string
	}
	handlerMap map[string]handler
)

const (
	expireSecond = 1 * 60
)

var (
	hdm handlerMap
)

func init() {
	hds := []handler{&echoCommand{}, &helpCommand{}}
	hdm = handlerMap{}
	for _, hd := range hds {
		hdm[hd.Key()] = hd
	}
}

func handle(msg *WxMessage) error {
	cs := strings.Split(string(msg.Content), " ")
	msg.Command = cs[0]
	msg.Args = cs[1:]

	step, err := redis.String(redisPool.Get().Do("GET", fmt.Sprintf("%s-%s", msg.FromUserName, msg.Command)))
	if err != nil && err != redis.ErrNil {
		log.Println("Redis GET error:", err.Error())
		return err
	}
	if err == nil {
		msg.Step = step
		return hdm[msg.Command].Do(msg)
	}

	if hd, ok := hdm[string(msg.Command)]; ok {
		return hd.Do(msg)
	}
	return robot(msg)
}
