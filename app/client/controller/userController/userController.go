package userController

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	userDao3 "kChatRoom/app/client/dao/userDao"
	userModel2 "kChatRoom/app/client/model/userModel"
	"kChatRoom/common/global"
	"kChatRoom/common/message"
	"kChatRoom/utils/help"
	"net/http"
	"strings"
	"time"
)

//Login user login
func Login(ctx *gin.Context) {
	mail := strings.TrimSpace(ctx.Query("mail"))
	pwd := strings.TrimSpace(ctx.Query("password"))
	userDao := userDao3.NewUserDao()
	msg := &message.RequestMsg{}
	findUser := userDao.GetUserByMail(mail)
	fmt.Println(findUser)
	if findUser.ID > 0 {
		md5Pwd := fmt.Sprintf("%x", md5.Sum([]byte(pwd)))
		fmt.Println(md5Pwd)
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
	mail := strings.TrimSpace(ctx.Query("mail"))
	userName := strings.TrimSpace(ctx.Query("username"))
	password := fmt.Sprintf("%x", md5.Sum([]byte(strings.TrimSpace(ctx.Query("password")))))
	verCode := strings.TrimSpace(ctx.Query("vercode"))
	//校验验证码
	rd := global.RedisPoolGlobal.Get()
	defer rd.Close()
	msg := &message.RequestMsg{}

	trueCode, _ := redis.String(rd.Do("Get", mail))

	if trueCode == verCode {
		user := &userModel2.UserModel{
			UserName: userName,
			Mail:     mail,
			Password: password,
		}
		userDao := userDao3.NewUserDao()

		res, m := userDao.AddUser(user)
		if res {
			rd.Do("Del", mail)
			msg.Code = 200
			msg.Msg = "注册成功"
		} else {
			msg.Code = 300
			msg.Msg = m
		}
	} else {
		msg.Code = 300
		msg.Msg = "验证码错误或失效！"
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
	err := help.SendMail([]string{mail}, "kChatRoom注册验证码", fmt.Sprintf("<h1>kChatRoom注册验证码为：%s</h1><br><p>15分钟内有效</p>", code))
	if err != nil {
		msg.Code = 300
		msg.Msg = fmt.Sprintf("验证码发送失败！,%s", err)
	} else {
		msg.Code = 200
		msg.Msg = "发送成功！"
	}

	ctx.JSON(http.StatusOK, msg)
}
