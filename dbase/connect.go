package dbase

import (
	"fiber_jwt/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "root:root123456@tcp(127.0.0.1:3306)/fiber_jwt?charset=utf8&parseTime=true&loc=Local"
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	} else {
		log.Println("连接数据库成功")
	}

	DB = connection

	connection.AutoMigrate(&models.Users{})
}
