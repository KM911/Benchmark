package dao

import (
	"app/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func InitMySql() {
	// 连接数据库
	DB, err = gorm.Open(mysql.Open(config.KM911LocalDSN), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// 显示 SQL 语句
}
