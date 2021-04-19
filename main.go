package main

import (
	"kChatRoom/common/global"
	"kChatRoom/server"
)

func init() {
	global.CfgInit()

}

func main() {
	err := server.SetupRouter().Run(":8060")
	if err != nil {
		return
	}
}
