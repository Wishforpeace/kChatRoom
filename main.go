package main

import "y2/kChatRoom/server"

func init() {

}

func main() {
	err := server.SetupRouter().Run(":8060")
	if err != nil {
		return
	}
}
