package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	userDao2 "kChatRoom/app/client/dao/userDao"
	"kChatRoom/app/service/model"
	"kChatRoom/common/global"
	"kChatRoom/common/message"
)

// Process 处理连接
func Process(conn *websocket.Conn, mail string) {

	userDao := userDao2.NewUserDao()
	user := userDao.GetUserByMail(mail)

	clint := &model.Client{
		Conn: conn,
		User: user,
	}
	if _, ok := global.ClientsGlobal[user.Mail]; !ok {
		global.ClientsGlobal[user.Mail] = clint
	}

	//加入成功消息
	global.OnlineChan <- clint

	//失去连接时处理的事情
	defer func() {
		//去除在线状态
		if _, ok := global.ClientsGlobal[user.Mail]; ok {
			delete(global.ClientsGlobal, user.Mail)
		}
		//添加用户离开消息
		global.LeaveChan <- clint
		//关闭链接
		conn.Close()
	}()
	for {
		//一直读取消息
		_, msgStr, err := conn.ReadMessage()
		if err != nil {
			break
		}
		//fmt.Println("ws:", string(msgStr))
		//处理消息
		ProcessMessage(msgStr)
	}
}

//ProcessMessage 处理接受消息
func ProcessMessage(msg []byte) {
	Msg := &message.Message{}
	//fmt.Println(string(msg))
	err := json.Unmarshal(msg, Msg)
	if err != nil {
		fmt.Println("unmarshal err", err)
		return
	}
	global.MessageChan <- Msg
}
