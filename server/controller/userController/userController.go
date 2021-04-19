package userController

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"kChatRoom/common/global"
	"kChatRoom/common/message"
	userDao2 "kChatRoom/server/dao/userDao"
	"kChatRoom/server/model/userModel"
	"kChatRoom/utils/help"
	"net/http"
	"strings"
	"time"
)

//Login user login
func Login(ctx *gin.Context) {
	mail := strings.TrimSpace(ctx.PostForm("mail"))
	pwd := strings.TrimSpace(ctx.PostForm("password"))
	userDao := userDao2.NewUserDao()
	msg := &message.RequestMsg{}
	findUser := userDao.GetUserByMail(mail)
	if findUser.ID != 0 {
		md5Pwd := fmt.Sprintf("%x", md5.Sum([]byte(pwd)))
		if md5Pwd == findUser.Password {
			help.SetCookie("user", findUser.Mail, ctx)
			msg.Code = 200
			msg.Msg = "登陆成功！"
		} else {
			msg.Code = 300
			msg.Msg = "用户名或密码错误！"
		}
	} else {
		msg.Code = 300
		msg.Msg = "用户名或密码错误！"
	}
	ctx.JSON(http.StatusOK, msg)
}

// Register register
func Register(ctx *gin.Context) {
	userName := strings.TrimSpace(ctx.Query("username"))
	password := fmt.Sprintf("%x", md5.Sum([]byte(strings.TrimSpace(ctx.Query("password")))))
	//verCode  :=  strings.TrimSpace(ctx.Query("vercode"))

	//校验验证码

	user := &userModel.UserModel{
		UserName: userName,
		Password: password,
	}
	userDao := userDao2.NewUserDao()
	msg := &message.RequestMsg{}
	res, m := userDao.AddUser(user)
	if res {
		msg.Code = 200
		msg.Msg = "注册成功"
	} else {
		msg.Code = 300
		msg.Msg = m
	}
	ctx.JSON(http.StatusOK, msg)
}

// SendVerCode 发送验证码
func SendVerCode(ctx *gin.Context) {
	mail := strings.TrimSpace(ctx.Query("mail"))
	rd := global.RedisPoolGlobal.Get()
	defer rd.Close()
	msg := &message.RequestMsg{}

	code := help.CreateValidateCode(6)
	_, _ = rd.Do("Set", mail, code)
	_, _ = rd.Do("EXPIRE", mail, 30*time.Minute)
	err := help.SendMail([]string{"738256016@qq.com"}, "kChatRoom注册验证码", fmt.Sprintf("<h1>kChatRoom注册验证码为：%s</h1><br><p>15分钟内有效</p>", code))
	if err != nil {
		msg.Code = 300
		msg.Msg = "发送验证码失败！"
	} else {
		msg.Code = 200
		msg.Msg = "发送成功！"
	}

	ctx.JSON(http.StatusOK, msg)
}
