package app

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	userController2 "kChatRoom/app/client/controller/userController"
	"kChatRoom/app/service"
	"kChatRoom/common/global"
	"kChatRoom/common/message"
	"kChatRoom/utils/help"
	"net/http"
)

//LoginAuth 登录后刷新权限组缓存
func LoginAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		white := map[string]string{
			"/view/login":          "/view/login",          //登陆视图
			"/view/login-action":   "/view/login-action",   //登陆动作
			"/view/login/sendCode": "/view/login/sendCode", //发送验证码
			"/view/register":       "/view/register",       //注册
			"/view/logout":         "/view/logout",         //退出登陆
		}
		nowUrl := c.Request.URL.Path
		if _, ok := white[nowUrl]; ok != true {
			nowMail, err := c.Cookie("user")
			nowAuth, err := c.Cookie("auth")
			if err != nil || nowMail == "" || nowAuth == "" {
				c.Redirect(http.StatusMovedPermanently, "/view/login")
			}
		}
		c.Next()
	}
}

//CheckLogin 校验是否登陆
func CheckLogin(c *gin.Context) {
	white := map[string]string{
		"/view/login":          "/view/login",          //登陆视图
		"/view/login-action":   "/view/login-action",   //登陆动作
		"/view/login/sendCode": "/view/login/sendCode", //发送验证码
		"/view/register":       "/view/register",       //注册
	}
	nowUrl := c.Request.URL.Path
	if _, ok := white[nowUrl]; ok != true {
		nowMail, err := c.Cookie("user")
		nowAuth, err := c.Cookie("auth")
		if err != nil || nowMail == "" || nowAuth == "" {
			fmt.Println("登陆过期！")
			c.Redirect(http.StatusMovedPermanently, "/view/login")
		}
	}
}

// SetupRouter router
func SetupRouter() *gin.Engine {
	r := gin.New()

	//引入视图/静态资源
	r.LoadHTMLGlob("app/client/views/**/*")
	r.Static("/static", "./static/service")

	//首页跳转
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/view/login")
	})

	r.GET("test", func(c *gin.Context) {
		rd := global.RedisPoolGlobal.Get()
		defer rd.Close()

		return
		//help.SetCookie("test","123",3600,c)
		code := help.CreateValidateCode(6)
		help.SendMail([]string{"738256016@qq.com"}, "kChatRoom注册验证码", fmt.Sprintf("<h1>kChatRoom注册验证码为：%s</h1><br><p>15分钟内有效</p>", code))
	})

	//前台视图
	view := r.Group("view")
	{
		//用户登陆/注册
		{
			view.GET("login", func(c *gin.Context) {
				CheckLogin(c)
				user, _ := c.Cookie("user")
				auth, _ := c.Cookie("auth")
				if user != "" && auth != "" {
					c.Redirect(http.StatusMovedPermanently, "/view/index")
				}
				c.HTML(http.StatusOK, "login.html", nil)
			})
			//提交登陆
			view.GET("login-action", userController2.Login)
			//验证码
			view.GET("login/sendCode", userController2.SendVerCode)
			//注册
			view.GET("register", userController2.Register)
			//退出登陆
			view.GET("logout", func(c *gin.Context) {
				help.DelCookie("user", c)
				help.DelCookie("auth", c)
				CheckLogin(c)
			})

			view.GET("test", func(c *gin.Context) {
				c.String(http.StatusOK, "hello")
			})
		}
		//主界面
		{
			view.GET("index", func(c *gin.Context) {
				CheckLogin(c)
				mail, _ := c.Cookie("user")
				key, _ := c.Cookie("auth")
				c.HTML(http.StatusOK, "chat.html", gin.H{
					"mail": mail,
					"key":  key,
				})
			})
		}

	}
	// 聊天请求
	chat := r.Group("service")
	{
		chat.GET("pong", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

		chat.GET("ws", service.Ws)
	}
	//api 接口
	api := r.Group("api")
	{

		api.GET("test", func(c *gin.Context) {
			msg := message.Message{
				Type:  message.MsgTypeLogin,
				Mail:  "wew@qq.com",
				Msg:   "test",
				ToUid: 10,
			}
			str, _ := json.Marshal(msg)
			fmt.Println(string(str))
		})
	}

	return r
}
