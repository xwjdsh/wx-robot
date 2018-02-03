package handle

import (
	"log"
	"time"

	"github.com/muesli/cache2go"
	"github.com/xwjdsh/wx-robot/utils"
)

type Echo struct{}

func (e *Echo) Do(msg *utils.WxMessage) error {
	log.Println("In echo method")
	cache := cache2go.Cache("cache")
	if msg.Step == "1" {
		if _, err := cache.Delete(msg.Key); err != nil {
			log.Println("cache error:", err.Error())
		}
	} else {
		msg.Content = "I will repeat your next words."
		cache.Add(msg.Key, 60*time.Second, "1")
	}
	msg.Reverse()
	return nil
}

func (*Echo) Desc() string {
	return "Repeat your next words"
}

func (*Echo) Key() string {
	return "@echo"
}
