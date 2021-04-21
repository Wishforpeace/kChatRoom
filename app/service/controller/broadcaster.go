package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"kChatRoom/common/global"
	"kChatRoom/common/message"
)

// Broadcaster 监听消息 消息转发
func Broadcaster() {
	for {
		select {
		//上线通知
		case client := <-global.OnlineChan:
			fmt.Printf("%s 加入了聊天室！", client.User.UserName)
			msg := &message.Message{
				Type: message.MsgTypeOnline,
				Mail: client.User.Mail,
				Name: client.User.UserName,
				Msg:  fmt.Sprintf("%s 加入了聊天室！", client.User.UserName),
			}
			global.MessageChan <- msg
		//离线通知
		case client := <-global.LeaveChan:
			fmt.Printf("%s 离开了聊天室！", client.User.UserName)
			msg := &message.Message{
				Type: message.MsgTypeLeave,
				Mail: client.User.Mail,
				Name: client.User.UserName,
				Msg:  fmt.Sprintf("%s 离开了聊天室！", client.User.UserName),
			}
			global.MessageChan <- msg
		//转发消息
		case msg := <-global.MessageChan:
			SendMsg(msg)
		}
	}
}

//SendMsg 处理转发消息
func SendMsg(msg *message.Message) {
	msgStr, _ := json.Marshal(msg)
	switch msg.Type {
	case message.MsgTypeLeave, message.MsgTypeSms, message.MsgTypeOnline:
		for _, client := range global.ClientsGlobal {
			//排除自己
			if client.User.Mail != msg.Mail {
				//if true {
				err := client.Conn.WriteMessage(websocket.TextMessage, msgStr)
				if err != nil {
					fmt.Println("send msg err:", err)
					return
				}
			}
		}
	//私发消息
	case message.MsgTypeSmsOne:

	}

}
