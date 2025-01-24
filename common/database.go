package common

import (
	"easydemo/model"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	host := "localhost"
	port := "3306"
	database := "easydb"
	username := "root"
	password := "rootpwd"
	charset := "utf8"

	// 构建DSN（Data Source Name）
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		username, password, host, port, database, charset)

	// 初始化 GORM DB
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database, err:" + err.Error())
	}
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		panic("failed to migrate database, err:" + err.Error())
	}
}

func GetDB() *gorm.DB {
	return db
}
