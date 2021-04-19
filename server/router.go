package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kChatRoom/utils/help"
	"net/http"
)

// SetupRouter router
func SetupRouter() *gin.Engine {
	r := gin.New()

	//引入视图/静态资源
	r.LoadHTMLGlob("server/views/**/*")
	r.Static("/static", "./static/chat")

	//首页跳转
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/view/login")
	})

	r.GET("test", func(c *gin.Context) {
		//help.SetCookie("test","123",3600,c)
		code := help.CreateValidateCode(6)
		help.SendMail([]string{"738256016@qq.com"}, "kChatRoom注册验证码", fmt.Sprintf("<h1>kChatRoom注册验证码为：%s</h1><br><p>15分钟内有效</p>", code))
	})

	//前台视图
	view := r.Group("view")
	{
		view.GET("login", func(c *gin.Context) {
			c.HTML(http.StatusOK, "login.html", nil)
		})

		view.GET("test", func(c *gin.Context) {
			c.String(http.StatusOK, "hello")
		})

	}
	// 聊天接口
	chat := r.Group("chat")
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
