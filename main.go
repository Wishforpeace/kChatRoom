package main

import (
	"kChatRoom/app"
	"kChatRoom/common/global"
)

func init() {
	global.CfgInit()
	global.GblInit()
}

func main() {
	err := app.SetupRouter().Run(":8060")
	if err != nil {
		return
	}
}
