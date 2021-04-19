package common

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

//ViperGlobal viper配置文件
var ViperGlobal *viper.Viper

//RedisPoolGlobal redis pool
var RedisPoolGlobal *redis.Pool

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
			if _, err := c.Do("AUTH", pwd); err != nil {
				err := c.Close()
				if err != nil {
					return nil, err
				}
				return nil, err
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
