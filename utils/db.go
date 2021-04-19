package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"kChatRoom/common"
)

var globalDb *gorm.DB

// NewDB gorm DB
func NewDB() *gorm.DB {
	if globalDb == nil {
		globalDb, _ = initMysql(common.ViperGlobal.Sub("database"))
	}
	sqlDb, _ := globalDb.DB()
	if err := sqlDb.Ping(); err != nil {
		_ = sqlDb.Close()
		globalDb, _ = initMysql(common.ViperGlobal.Sub("database"))
	}
	return globalDb
}

//InitMysql
func initMysql(cfg *viper.Viper) (*gorm.DB, error) {

	dsn := cfg.GetString("dsn")
	skipDefaultTransaction := cfg.GetBool("skipDefaultTransaction")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: skipDefaultTransaction,
	})
	if err != nil {
		fmt.Println(fmt.Sprintf("mysql conner err =%s", err))
		return nil, err
	}
	return db, nil
}
