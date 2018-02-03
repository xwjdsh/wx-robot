package handle

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/xwjdsh/wx-robot/utils"
)

const (
	URL  = "http://www.tuling123.com/openapi/api"
	TYPE = "application/json;charset=utf-8"
)

type robotRequest struct {
	Info   string `json:"info"`
	Key    string `json:"key"`
	Loc    string `json:"loc"`
	UserID string `json:"userid"`
}

type robotResponse struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

type Robot struct{}

func (r *Robot) Do(msg *utils.WxMessage) error {
	log.Println("In robot method")
	robotReq := &robotRequest{
		Info:   string(msg.Content),
		Key:    *utils.RobotFlag,
		UserID: string(msg.FromUserName),
	}
	robotResp, err := send(robotReq)
	if err != nil {
		return err
	}
	msg.Reverse()
	msg.MsgType = utils.CData("text")
	msg.Content = utils.CData(robotResp.Text)
	return nil
}

func (*Robot) Desc() string {
	return ""
}

func (*Robot) Key() string {
	return ""
}

func send(robotReq *robotRequest) (*robotResponse, error) {
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
	robotResp := &robotResponse{}
	if err := json.Unmarshal(result, robotResp); err != nil {
		return nil, err
	}
	return robotResp, nil
}
