package controller

import (
	"fmt"
	"kChatRoom/app/client/model/userModel"
	"kChatRoom/app/service/model"
	"kChatRoom/common/global"
	"kChatRoom/common/message"
	"time"
)

const (
	RobotName = "实习机器人"
	RobotMail = "robot@qq.com"
)

//InitRobot 初始化机器人
func InitRobot() {
	user := &userModel.UserModel{
		Mail:     RobotMail,
		UserName: RobotName,
	}
	botClient := &model.Client{
		Type: model.ClientTypeRobot,
		User: user,
	}
	if _, ok := global.ClientsGlobal[user.Mail]; !ok {
		global.ClientsGlobal[user.Mail] = botClient
	}
}

//WelcomeMsg 进入房间欢迎语
func WelcomeMsg(name string) {
	msg := &message.Message{
		Type: message.MsgTypeRobot,
		Mail: RobotMail,
		Name: RobotName,
		Msg:  fmt.Sprintf("欢迎 %s 加入聊天室~", name),
		Head: "",
	}
	go func() {
		time.Sleep(time.Second * 2)
		global.MessageChan <- msg
	}()

}
