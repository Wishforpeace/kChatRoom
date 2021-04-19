package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// SetupRouter router
func SetupRouter() *gin.Engine {
	r := gin.New()

	//引入视图/静态资源
	//r.LoadHTMLGlob("server/views/**/*")
	r.Static("/static", "./static/chat")
	//前台视图
	view := r.Group("view")
	{
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
