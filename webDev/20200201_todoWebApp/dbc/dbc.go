package dbc

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	// Db 匯出變數
	Db *gorm.DB
)

// InitDB 初始化DB
func InitDB() (err error) {
	dsn := "host=localhost user=postgres password=abc123 dbname=postgres port=1036 sslmode=disable TimeZone=Asia/Shanghai"
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return
}
