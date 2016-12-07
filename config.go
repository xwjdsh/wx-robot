package main

import (
	"encoding/json"
	"io/ioutil"
)

var (
	conf config
)

type config struct {
	Port        int    `json:"port"`
	RobotKey    string `json:"robotKey"`
	InfoLogPath string `json:"infoLogPath"`
	ErrLogPath  string `json:"errLogPath"`
	RedisConn   string `json:"redisConn"`
}

func init() {
	file, err := ioutil.ReadFile("setting.json")
	if err != nil {
		panic("Make sure `setting.json` file exists!")
	}

	err = json.Unmarshal(file, &conf)
	if err != nil {
		panic("Make sure config file correct!")
	}
}
