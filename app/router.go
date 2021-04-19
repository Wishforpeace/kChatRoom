package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	userController2 "kChatRoom/app/client/controller/userController"
	"kChatRoom/common/global"
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
		}
		nowUrl := c.Request.URL.Path
		if _, ok := white[nowUrl]; ok != true {
			nowMail, err := c.Cookie("user")
			if err != nil || nowMail == "" {
				c.Redirect(http.StatusMovedPermanently, "/view/login")
			}
		}
		c.Next()
	}
}

// SetupRouter router
func SetupRouter() *gin.Engine {
	r := gin.New()

	//引入视图/静态资源
	r.LoadHTMLGlob("app/client/views/**/*")
	r.Static("/static", "./static/server")

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
		//登陆校验
		view.Use(LoginAuth())
		//用户登陆/注册
		{
			view.GET("login", func(c *gin.Context) {
				c.HTML(http.StatusOK, "login.html", nil)
			})
			//提交登陆
			view.GET("login-action", userController2.Login)
			//验证码
			view.GET("login/sendCode", userController2.SendVerCode)
			//注册
			view.GET("register", userController2.Register)

			view.GET("test", func(c *gin.Context) {
				c.String(http.StatusOK, "hello")
			})
		}
		//主界面
		{
			view.GET("index", func(c *gin.Context) {
				mail, _ := c.Cookie("user")
				c.HTML(http.StatusOK, "chat.html", gin.H{
					"mail": mail,
				})
			})
		}

	}
	// 聊天接口
	chat := r.Group("server")
	{
		chat.GET("test", func(c *gin.Context) {

		})
	}
	//api 接口
	api := r.Group("api")
	{

		api.GET("test", func(c *gin.Context) {

		})
	}

	return r
}
