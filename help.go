package main

import (
	"fmt"
	"log"
)

type helpCommand struct{}

func (*helpCommand) Do(msg *WxMessage) error {
	log.Println("In help method")
	msg.Reverse()
	var content string
	for key, hd := range hdm {
		content += fmt.Sprintf("%s -- %s\n", key, hd.Desc())
	}
	msg.Content = cdata(content)
	return nil
}

func (*helpCommand) Desc() string {
	return "Display available commands"
}

func (*helpCommand) Key() string {
	return "@help"
}
