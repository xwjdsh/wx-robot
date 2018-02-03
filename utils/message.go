package utils

import (
	"encoding/xml"
	"time"
)

type WxMessage struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CData    `xml:"ToUserName"`
	FromUserName CData    `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      CData    `xml:"MsgType"`
	Content      CData    `xml:"Content"`
	MsgID        string   `xml:"MsgId"`
	Key          string   `xml:"-"`
	Command      string   `xml:"-"`
	Args         []string `xml:"-"`
	Step         string   `xml:"-"`
}

func (msg *WxMessage) Reverse() {
	msg.ToUserName, msg.FromUserName = msg.FromUserName, msg.ToUserName
	msg.CreateTime = time.Now().Unix()
}
