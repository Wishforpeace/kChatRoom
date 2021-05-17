package global

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
	"io/ioutil"
	"kChatRoom/app/service/model"
	"kChatRoom/common/message"
	"kChatRoom/utils/cookie"
	"kChatRoom/utils/mail"
	"log"
	"os"
	"strings"
	"time"
)

//ViperGlobal viper配置文件
var ViperGlobal *viper.Viper

//RedisPoolGlobal redis pool
var RedisPoolGlobal *redis.Pool

// CookieGlobal  cookie config
var CookieGlobal *cookie.Cookie

//MailGlobal mail.Mail
var MailGlobal *mail.Mail

//ClientsGlobal 在线用户列表
var ClientsGlobal map[string]*model.Client

//OnlineChan 在线用户加入
var OnlineChan chan *model.Client

//LeaveChan 离线用户加入
var LeaveChan chan *model.Client

//MessageChan 等待发送消息加入
var MessageChan chan *message.Message

// LoginUsers 二次验证过登陆的用户
var LoginUsers map[string]bool

//GblInit 初始化
func GblInit() {
	ClientsGlobal = make(map[string]*model.Client)
	OnlineChan = make(chan *model.Client, 10)
	LeaveChan = make(chan *model.Client, 10)
	MessageChan = make(chan *message.Message, 10)
	LoginUsers = make(map[string]bool)
}

// CfgInit 载入配置文件
func CfgInit() {
	path := "config/config.yml"
	viper.SetConfigFile(path)
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(fmt.Sprintf("Read config file fail: %s", err.Error()))
	}

	//Replace environment variables
	err = viper.ReadConfig(strings.NewReader(os.ExpandEnv(string(content))))
	if err != nil {
		log.Fatal(fmt.Sprintf("Parse config file fail: %s", err.Error()))
	}

	//global viper
	ViperGlobal = viper.Sub("settings")

	cfgRedis := viper.Sub("settings.redis")
	initRedis(cfgRedis)

	cookieRedis := viper.Sub("settings.cookie")
	InitCookie(cookieRedis)

	mailCfg := viper.Sub("mail")
	InitMail(mailCfg)

}

//initRedis 初始化redis
func initRedis(cfg *viper.Viper) {
	pool := &redis.Pool{
		MaxIdle:     cfg.GetInt("maxIdle"),                    //最大空闲连接数  8
		MaxActive:   cfg.GetInt("maxActive"),                  //表示和数据库大最大链接数，0表示不限制
		IdleTimeout: time.Duration(cfg.GetInt("idleTimeout")), //最大空闲时间 100
		Dial: func() (redis.Conn, error) { //初始化链接的代码
			pwd := cfg.GetString("password")
			c, err := redis.Dial(cfg.GetString("netWork"), cfg.GetString("address"))
			if pwd != "" {
				if _, err := c.Do("AUTH", pwd); err != nil {
					err := c.Close()
					if err != nil {
						return nil, err
					}
					return nil, err
				}
			}
			db := cfg.GetInt("dbSelect")
			if _, err := c.Do("SELECT", db); err != nil {
				err := c.Close()
				if err != nil {
					return nil, err
				}
				return nil, err
			}
			return c, err
		},
	}
	RedisPoolGlobal = pool
}

// InitCookie cookie
func InitCookie(cfg *viper.Viper) {
	CookieGlobal = &cookie.Cookie{
		Path:     cfg.GetString("path"),
		Domain:   cfg.GetString("cookieDomain"),
		Secure:   cfg.GetBool("secure"),
		HttpOnly: cfg.GetBool("httpOnly"),
	}
}

//InitMail mail
func InitMail(cfg *viper.Viper) {
	MailGlobal = &mail.Mail{
		User: cfg.GetString("user"),
		Pass: cfg.GetString("pass"),
		Host: cfg.GetString("host"),
		Port: cfg.GetString("port"),
	}
}
