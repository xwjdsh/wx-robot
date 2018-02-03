package handle

import (
	"fmt"
	"log"

	"github.com/xwjdsh/wx-robot/utils"
)

type Help struct{}

func (h *Help) Do(msg *utils.WxMessage) error {
	log.Println("In help method")
	msg.Reverse()
	echo := &Echo{}
	content := fmt.Sprintf("%s -- %s\n", echo.Key(), echo.Desc())
	content += fmt.Sprintf("%s -- %s\n", h.Key(), h.Desc())
	msg.Content = utils.CData(content)
	return nil
}

func (*Help) Desc() string {
	return "Display available commands"
}

func (*Help) Key() string {
	return "@help"
}
