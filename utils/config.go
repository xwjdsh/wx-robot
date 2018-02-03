package utils

import (
	"flag"
	"os"
)

var (
	PortFlag  = flag.Int("p", 3000, "port")
	RobotFlag = flag.String("robot", "", "robot key")
)

func InitConfig() {
	flag.Parse()
	if robot := os.Getenv("ROBOT"); robot != "" {
		RobotFlag = &robot
	}
}
