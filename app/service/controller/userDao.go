package controller

import (
	"github.com/gin-gonic/gin"
	userDao2 "kChatRoom/app/client/dao/userDao"
	"kChatRoom/common/global"
	"kChatRoom/common/message"
	"net/http"
)

// SaveHead 保存用户头像
func SaveHead(c *gin.Context) {
	mail, err := c.Cookie("user")
	config := c.Query("headConfig")
	Msg := &message.RequestMsg{}
	if err != nil || mail == "" || config == "" {
		Msg.Code = 100
		Msg.Msg = "保存失败！"
	} else {
		userDao := userDao2.NewUserDao()
		userDao.SaveHead(mail, config)
		Msg.Code = 200
		Msg.Msg = "保存成功！"
	}
	c.JSON(http.StatusOK, Msg)
}

//GetOnlineUser 获取在线用户人数
func GetOnlineUser(c *gin.Context) {
	Msg := &message.RequestMsg{}
	onlineNum := len(global.ClientsGlobal)
	Msg.Res = onlineNum
	Msg.Code = 200
	Msg.Msg = "保存成功！"
	c.JSON(http.StatusOK, Msg)
}

// GetUserInfo 获取用户基本信息
func GetUserInfo(c *gin.Context) {
	msg := message.RequestMsg{Code: http.StatusOK}
	mail := c.Query("mail")
	userDao := userDao2.NewUserDao()
	user := userDao.GetUserByMail(mail)
	user.Password = ""
	msg.Res = user
	c.JSON(http.StatusOK, msg)
}
