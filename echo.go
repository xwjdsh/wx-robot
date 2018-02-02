package main

import "log"

type echoCommand struct{}

func (e *echoCommand) Do(msg *WxMessage) error {
	log.Println("In echo method")
	if msg.Step == "1" {
		redisPool.Get().Do("DEL", msg.Key)
	} else {
		msg.Content = "I will repeat your next words."
		redisPool.Get().Do("SETEX", msg.Key, expireSecond, "1")
	}
	msg.Reverse()
	return nil
}

func (*echoCommand) Desc() string {
	return "Repeat your next words"
}

func (*echoCommand) Key() string {
	return "@echo"
}
