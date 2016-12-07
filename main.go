package main

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/iris-contrib/middleware/logger"
	"github.com/kataras/iris"
	"github.com/op/go-logging"
)

var (
	logObj         = logging.MustGetLogger("wx-robot")
	errorLogFormat = logging.MustStringFormatter(
		`%{time:15:04:05} %{shortfunc} ▶ %{level:.4s} %{id:03x} %{message}`,
	)
	infoLogFormat = logging.MustStringFormatter(
		`%{time:15:04:05} ▶ %{level:.4s} %{message}`,
	)
)

func main() {
	initLog()
	initRedis()
	api := iris.New()
	api.Use(logger.New())
	api.Post("/", getMsg)
	api.Listen(fmt.Sprintf(":%d", conf.Port))
}

func initLog() {
	getFile := func(filePath string) *os.File {
		if len(filePath) == 0 {
			return nil
		}
		_, err := os.Stat(filePath)
		// create file if not exists
		if os.IsNotExist(err) {
			if _, err = os.Create(filePath); err != nil {
				return nil
			}
		}
		f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 600)
		if err != nil {
			return nil
		}
		return f
	}

	out := os.Stdout
	if f := getFile(conf.InfoLogPath); f != nil {
		out = f
	}
	backend1 := logging.NewLogBackend(out, "", 600)
	backend1Formatter := logging.NewBackendFormatter(backend1, infoLogFormat)

	out = os.Stderr
	if f := getFile(conf.ErrLogPath); f != nil {
		out = f
	}
	backend2 := logging.NewLogBackend(out, "", 600)
	backend2Formatter := logging.NewBackendFormatter(backend2, errorLogFormat)
	backend2Leveled := logging.AddModuleLevel(backend2Formatter)
	backend2Leveled.SetLevel(logging.ERROR, "")

	logging.SetBackend(backend1Formatter, backend2Leveled)
}

func getMsg(c *iris.Context) {
	logObj.Info(fmt.Sprintf("Message in:\n%s\n", string(c.RequestCtx.Request.Body())))
	msg := &WxMessage{}
	if err := c.ReadXML(msg); err != nil {
		logObj.Info("Read xml error:", err.Error())
		return
	}
	err := handle(msg)
	if err != nil {
		logObj.Error("Handle message error:", err.Error())
		c.Write("")
		return
	}
	data, _ := xml.MarshalIndent(msg, "", "\t")
	logObj.Info(fmt.Sprintf("Response:\n%s\n", string(data)))
	c.XML(iris.StatusOK, msg)
}
