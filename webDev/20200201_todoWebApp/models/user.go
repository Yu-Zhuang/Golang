package models

import (
	"todoWebApp/dbc"

	"gorm.io/gorm"
)

// User 使用者
type User struct {
	gorm.Model
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Todo     []Todo `gorm:"foreignKey:UserRefer"`
}

// CreateUser 新增人
func CreateUser(user *User) (err error) {
	result := dbc.Db.Create(&user)
	return result.Error
}
