package main

import (
	"kChatRoom/app"
	"kChatRoom/app/service/controller"
	"kChatRoom/common/global"
)

func init() {
	global.CfgInit()
	global.GblInit()
	//消息广播
	go controller.Broadcaster()
	//聊天机器人
	controller.InitRobot()
}

func main() {
	err := app.SetupRouter().Run(":8060")
	if err != nil {
		return
	}
}
