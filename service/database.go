package service

import (
	//"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dbConn *gorm.DB

func CreateConnect() {
	//dsn := "root:@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	dbConn = db
}

func GetDatabaseConnection() *gorm.DB {
	sqlDB, err := dbConn.DB()
	if err != nil {
		return dbConn
	}
	if err := sqlDB.Ping(); err != nil {
		return dbConn
	}
	return dbConn
}
