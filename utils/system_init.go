package utils

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitMySQL() {
	//dsn := "root:kzl180@tcp(103.116.246.172:3306)/video_sharing?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:kzl180@tcp(127.0.0.1:3306)/video_sharing?charset=utf8mb4&parseTime=True&loc=Local"
	DB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
