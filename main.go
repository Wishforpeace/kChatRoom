package main

import (
	"kChatRoom/common"
	"kChatRoom/server"
)

func init() {
	common.CfgInit()
}

func main() {
	err := server.SetupRouter().Run(":8060")
	if err != nil {
		return
	}
}
