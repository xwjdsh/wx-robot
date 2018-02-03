package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/muesli/cache2go"
	"github.com/xwjdsh/wx-robot/handle"
	"github.com/xwjdsh/wx-robot/utils"
)

type (
	handler interface {
		Do(*utils.WxMessage) error
		Desc() string
		Key() string
	}
)

var (
	hdm   = map[string]handler{}
	cache = cache2go.Cache("cache")
)

func init() {
	hds := []handler{&handle.Echo{}, &handle.Help{}, &handle.Robot{}}
	for _, hd := range hds {
		hdm[hd.Key()] = hd
	}
}

func handleMsg(msg *utils.WxMessage) error {
	cs := strings.Split(string(msg.Content), " ")
	msg.Command = cs[0]
	msg.Args = cs[1:]

	item, err := cache2go.Cache("cache").Value(fmt.Sprintf("%s-%s", msg.FromUserName, msg.Command))
	if err != nil && err != cache2go.ErrKeyNotFound {
		log.Println("cache error:", err.Error())
		return err
	}
	if err == cache2go.ErrKeyNotFound {
		var err error
		if hd, ok := hdm[string(msg.Command)]; ok {
			err = hd.Do(msg)
		} else {
			err = hdm[""].Do(msg)
		}
		return err
	}
	msg.Step = item.Data().(string)
	return hdm[msg.Command].Do(msg)
}
