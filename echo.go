package main

type echo struct{}

func (e *echo) Do(msg *WxMessage) error {
	logObj.Info("In `echo` method")
	cc := string(msg.FromUserName) + currentCommand
	if msg.HasCommand {
		go redisPool.Get().Do("DEL", cc)
	} else {
		msg.Content = "I will repeat your next words."
		go redisPool.Get().Do("SETEX", cc, expireSecond, e.Key())
	}
	msg.Reverse()
	return nil
}

func (*echo) Desc() string {
	return "Repeat your next sentence"
}

func (*echo) Key() string {
	return "@echo"
}
