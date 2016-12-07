package main

import "github.com/garyburd/redigo/redis"

type (
	handler interface {
		Do(*WxMessage) error
		Desc() string
		Key() string
	}
	handlerMap map[string]handler
)

const (
	expireSecond   = 10 * 60
	currentCommand = "-Command"
	currentStep    = "-Step"
)

var (
	hdm handlerMap
)

func init() {
	hds := []handler{&echo{}, &help{}}
	hdm = handlerMap{}
	for _, hd := range hds {
		hdm[hd.Key()] = hd
	}
}

func handle(msg *WxMessage) error {
	cc, err := redis.String(redisPool.Get().Do("GET", string(msg.FromUserName)+currentCommand))
	if err != nil && err != redis.ErrNil {
		logObj.Error("Redis `GET` error:", err.Error())
	}
	if err == nil {
		msg.HasCommand = true
		return hdm[cc].Do(msg)
	}

	if hd, ok := hdm[string(msg.Content)]; ok {
		return hd.Do(msg)
	}
	return robot(msg)
}

func robot(msg *WxMessage) error {
	robotResp, err := send(string(msg.Content), string(msg.FromUserName))
	if err != nil {
		return err
	}
	msg.Reverse()
	msg.MsgType = cdata("text")
	msg.Content = cdata(robotResp.Text)
	return nil
}
