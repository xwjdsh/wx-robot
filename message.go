package main

import (
	"encoding/xml"
	"time"
)

type WxMessage struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   cdata    `xml:"ToUserName"`
	FromUserName cdata    `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      cdata    `xml:"MsgType"`
	Content      cdata    `xml:"Content"`
	MsgID        string   `xml:"MsgId"`
	HasCommand   bool     `xml:"-"`
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
