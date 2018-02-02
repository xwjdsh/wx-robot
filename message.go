package main

import (
	"encoding/xml"
	"time"
)

type cdata string

func (n cdata) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(struct {
		S string `xml:",innerxml"`
	}{
		S: "<![CDATA[" + string(n) + "]]>",
	}, start)
}

type WxMessage struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   cdata    `xml:"ToUserName"`
	FromUserName cdata    `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      cdata    `xml:"MsgType"`
	Content      cdata    `xml:"Content"`
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

type RobotRequest struct {
	Info   string `json:"info"`
	Key    string `json:"key"`
	Loc    string `json:"loc"`
	UserID string `json:"userid"`
}

type RobotResponse struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}
