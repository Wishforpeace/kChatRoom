package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/websocket"
	"kChatRoom/app/service/controller"
	"kChatRoom/common/global"
	"net/http"
)

//authCheck 安全登陆校验
func authCheck(key, mail string) bool {
	if key == "" || mail == "" {
		return false
	}
	//判断是否升级登陆
	_, ok := global.LoginUsers[mail]
	if ok {
		return true
	} else {
		rd := global.RedisPoolGlobal.Get()
		defer rd.Close()
		//判断redis登陆信息中是否存在 存在则升级登陆
		str, err := redis.String(rd.Do("Get", fmt.Sprintf("login_%s", mail)))
		//没有获取到数据
		if str == "" || err != nil {
			fmt.Println("非法key")
			return false
		} else if key == str { //升级登陆
			rd.Do("Del", fmt.Sprintf("login_%s", mail))
			global.LoginUsers[mail] = true
			return true
		} else {
			fmt.Println("非法key!")
			return false
		}
	}
}

// Ws 连接客户端
func Ws(c *gin.Context) {
	key := c.Query("key")
	mail := c.Query("mail")
	res := authCheck(key, mail)
	if !res {
		fmt.Println("非法登陆！")
		return
	}

	//升级协议 用户验证
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}
	//处理conn
	go controller.Process(conn, mail)
}
