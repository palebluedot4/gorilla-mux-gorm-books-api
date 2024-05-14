package config

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB() {
	mysqlDB, err := gorm.Open(mysql.Open("Bruce:palebluedot4@tcp(localhost:3306)/bookstoreAPI?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	db = mysqlDB
}

func GetDB() *gorm.DB {
	return db
}
