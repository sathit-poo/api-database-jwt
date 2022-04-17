package service

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbConn *gorm.DB

func CreateConnect() {
	//	dsn := "root:@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	//heroku
	dsn := "host=ec2-52-21-136-176.compute-1.amazonaws.com user=wwrxguwgoqqwbg password=54c7d62d01812dc04c0d0f25ba9ced413f1884214dbae2204b0620c72f1bd47c dbname=d7g6h8fm7bfvdt port=5432 sslmode=require TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
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
