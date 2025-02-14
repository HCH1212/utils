package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

func InitMysql(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	_ = db.AutoMigrate()

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)               // 最大空闲连接数
	sqlDB.SetMaxOpenConns(100)              // 最大可容纳
	sqlDB.SetConnMaxLifetime(time.Hour * 4) // 连接最大复用时间， 不能超过mysql的wait_timeout
	return db
}
