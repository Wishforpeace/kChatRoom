package help

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
	"kChatRoom/common/global"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

//SetCookie 设置基本cookie信息
func SetCookie(key, val string, c *gin.Context) {
	c.SetCookie(key, val, 86400, global.CookieGlobal.Path, global.CookieGlobal.Domain, global.CookieGlobal.Secure, global.CookieGlobal.HttpOnly)
}

//DelCookie 删除cookie
func DelCookie(key string, c *gin.Context) {
	c.SetCookie(key, "", -1, global.CookieGlobal.Path, global.CookieGlobal.Domain, global.CookieGlobal.Secure, global.CookieGlobal.HttpOnly)
}

//CreateValidateCode 生成验证码
//width 验证码长度
func CreateValidateCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}

//SendMail 发送邮件
func SendMail(mailTo []string, subject, body string) error {
	//mailConn := map[string]string{
	//  "user": "xxx@163.com",
	//  "pass": "your password",
	//  "host": "smtp.163.com",//smtp.exmail.qq.com smtp.qq.com
	//  "port": "465",
	//}
	//定义邮箱服务器连接信息，如果是网易邮箱 pass填密码，qq邮箱填授权码
	mailConn := map[string]string{
		"user": global.ViperGlobal.Sub("mail").GetString("user"),
		"pass": global.ViperGlobal.Sub("mail").GetString("pass"),
		"host": global.ViperGlobal.Sub("mail").GetString("host"),
		"port": global.ViperGlobal.Sub("mail").GetString("port"),
	}
	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int

	m := gomail.NewMessage()

	m.SetHeader("From", m.FormatAddress(mailConn["user"], "kChatRoom官方")) //这种方式可以添加别名，即“XX官方”
	//说明：如果是用网易邮箱账号发送，以下方法别名可以是中文，如果是qq企业邮箱，以下方法用中文别名，会报错，需要用上面此方法转码
	//m.SetHeader("From", "FB Sample"+"<"+mailConn["user"]+">") //这种方式可以添加别名，即“FB Sample”， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	//m.SetHeader("From", mailConn["user"])
	m.SetHeader("To", mailTo...)    //发送给多个用户
	m.SetHeader("Subject", subject) //设置邮件主题
	m.SetBody("text/html", body)    //设置邮件正文

	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])

	err := d.DialAndSend(m)
	return err

}
