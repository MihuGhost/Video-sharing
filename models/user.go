package models

import (
	"OnlineVideo/utils"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id        string `gorm:"primaryKey"`
	Name      string
	Password  string
	Phone     string
	Email     string
	Salt      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// 绑定表名
func (table *User) TableName() string {
	return "user"
}

func FindUserByName(name string) User {
	user := User{}
	utils.DB.Where("name = ?", name).First(&user)
	return user
}

func CreateUser(user User) *gorm.DB {
	return utils.DB.Create(&user)
}
