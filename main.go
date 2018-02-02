package main

import (
	"bytes"
	"encoding/xml"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

var (
	portFlag  = flag.Int("p", 3000, "port")
	robotFlag = flag.String("robot", "", "robot key")
	redisFlag = flag.String("redis", "localhost:6379", "redis connection")
)

func main() {
	initConfig()
	initRedis()
	e := echo.New()
	e.POST("/", getMsg)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", *portFlag)))
}

func getMsg(c echo.Context) error {
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request().Body)
	log.Printf("request: >>>>\n%s\n<<<<", buf.String())
	msg := &WxMessage{}
	if err := c.Bind(msg); err != nil {
		log.Println("Read xml error:", err.Error())
		return err
	}
	err := handle(msg)
	if err != nil {
		log.Println("Handle message error:", err.Error())
		return c.String(http.StatusOK, "")
	}
	data, _ := xml.MarshalIndent(msg, "", "\t")
	log.Printf("response: >>>>\n%s\n<<<<", string(data))
	return c.XML(http.StatusOK, msg)
}

func initConfig() {
	flag.Parse()
	if robot := os.Getenv("ROBOT"); robot != "" {
		robotFlag = &robot
	}
	if redisConn := os.Getenv("REDIS"); redisConn != "" {
		redisFlag = &redisConn
	}
}
