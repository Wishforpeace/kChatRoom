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

//onlyLoginCheck 唯一登陆
func onlyLoginCheck(client *model.Client) bool {
	msg := &message.Message{
		Type: message.MsgTypeOnlyLogin,
		Mail: client.User.Mail,
		Name: client.User.UserName,
		Msg:  "你的账户在其他设备登陆！",
		Head: client.User.Head,
	}

	rd := global.RedisPoolGlobal.Get()
	defer rd.Close()
	/*isOk := true

	//唯一登陆
	onlyStr, err := redis.String(rd.Do("HGet","login_only", client.User.Mail))
	if err != nil || onlyStr == ""{ //error
		isOk = false
	}else{
		onlyTime := common.Decrypt(onlyStr,[]byte("1d12jha8"))
		NewOnly, err := strconv.ParseInt(onlyTime, 10, 64)
		if err != nil || {
			isOk = false
		}
	}*/

	global.MessageChan <- msg
	return true
}

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
