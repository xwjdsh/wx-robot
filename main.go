package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/xwjdsh/wx-robot/utils"
)

func main() {
	utils.InitConfig()
	e := echo.New()
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	e.POST("/", getMsg)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", *utils.PortFlag)))
}

func getMsg(c echo.Context) error {
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request().Body)
	log.Printf("request: >>>>\n%s\n<<<<", buf.String())
	msg := &utils.WxMessage{}
	if err := c.Bind(msg); err != nil {
		log.Println("Read xml error:", err.Error())
		return err
	}
	err := handleMsg(msg)
	if err != nil {
		log.Println("Handle message error:", err.Error())
		return c.String(http.StatusOK, "")
	}
	data, _ := xml.MarshalIndent(msg, "", "\t")
	log.Printf("response: >>>>\n%s\n<<<<", string(data))
	return c.XML(http.StatusOK, msg)
}
