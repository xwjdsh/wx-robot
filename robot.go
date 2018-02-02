package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	URL  = "http://www.tuling123.com/openapi/api"
	TYPE = "application/json;charset=utf-8"
)

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

func send(message, userId string) (*RobotResponse, error) {
	robotReq := &RobotRequest{
		Info:   message,
		Key:    *robotFlag,
		UserID: userId,
	}
	b, _ := json.Marshal(robotReq)
	res, err := http.Post(URL, TYPE, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}
	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	robotResp := &RobotResponse{}
	if err := json.Unmarshal(result, robotResp); err != nil {
		return nil, err
	}
	return robotResp, nil
}
