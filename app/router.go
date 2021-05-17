package app

import (
	"github.com/gin-gonic/gin"
	userController2 "kChatRoom/app/client/controller/userController"
	"kChatRoom/app/client/dao/chatLogDao"
	"kChatRoom/app/service"
	"kChatRoom/app/service/controller"
	"kChatRoom/common/global"
	"kChatRoom/utils/help"
	"net/http"
	"strconv"
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
				c.Redirect(http.StatusFound, "/view/login")
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
			c.Redirect(http.StatusFound, "/view/login")
		}
	}
}

// SetupRouter router
func SetupRouter() *gin.Engine {
	r := gin.New()

	//引入视图/静态资源
	r.LoadHTMLGlob("app/client/views/**/*")
	r.Static("/static", "./static/chat")

	//首页跳转
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/view/login")
	})

	//前台视图
	view := r.Group("view")
	{
		//用户登陆/注册
		{
			view.GET("login", func(c *gin.Context) {
				user, _ := c.Cookie("user")
				auth, _ := c.Cookie("auth")
				if user != "" && auth != "" {
					c.Redirect(http.StatusFound, "/view/index")
					return
				}
				help.DelCookie("user", c)
				help.DelCookie("auth", c)
				c.HTML(http.StatusOK, "login.html", nil)
			})
			//提交登陆
			view.GET("login-action", userController2.Login)
			//验证码
			view.GET("login-sendCode", userController2.SendVerCode)
			//注册
			view.GET("register", userController2.Register)
			//退出登陆
			view.GET("logout", func(c *gin.Context) {
				c.SetCookie("user", "", -1, global.CookieGlobal.Path, global.CookieGlobal.Domain, global.CookieGlobal.Secure, global.CookieGlobal.HttpOnly)
				c.SetCookie("auth", "", -1, global.CookieGlobal.Path, global.CookieGlobal.Domain, global.CookieGlobal.Secure, global.CookieGlobal.HttpOnly)

				c.Redirect(http.StatusFound, "/view/login")
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

				url := "localhost:8060"
				if c.Request.Host != "127.0.0.1:8060" {
					url = c.Request.Host + ":8060"
				}

				c.HTML(http.StatusOK, "chatroom.html", gin.H{
					"mail": mail,
					"key":  key,
					"url":  url,
				})
			})
			view.GET("selectHead", func(c *gin.Context) {

				c.HTML(http.StatusOK, "select-head.html", gin.H{})
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
		api.GET("getUserInfo", controller.GetUserInfo)

		api.GET("saveHead", controller.SaveHead)

		api.GET("getOnlineUser", controller.GetOnlineUser)

		api.GET("rename", controller.Rename)

		api.GET("/getChatLog", func(c *gin.Context) {
			pageStr := c.DefaultQuery("page", "1")
			page, _ := strconv.Atoi(pageStr)
			limitStr := c.DefaultQuery("limit", "5")
			limit, _ := strconv.Atoi(limitStr)

			logDao := chatLogDao.NewChatLogDao()
			res := logDao.GetChatLog(page, limit)
			for _, v := range res {
				if v.Mail == controller.RobotMail {
					v.UserName = controller.RobotName
				}
			}
			c.JSON(http.StatusOK, res)
		})

		api.GET("test", func(c *gin.Context) {

		})
	}

	return r
}
