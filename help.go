package main

import "fmt"

type help struct{}

func (*help) Do(msg *WxMessage) error {
	logObj.Info("In `help` method")
	msg.Reverse()
	var content string
	for key, hd := range hdm {
		content += fmt.Sprintf("%s -- %s\n", key, hd.Desc())
	}
	msg.Content = cdata(content)
	return nil
}

func (*help) Desc() string {
	return "Display available interface"
}

func (*help) Key() string {
	return "@help"
}
