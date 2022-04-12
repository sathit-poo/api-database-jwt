package service

import (
	//	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbConn *gorm.DB

func CreateConnect() {
	//	dsn := "root:@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"

	//heroku
	dsn := "host=ec2-34-192-210-139.compute-1.amazonaws.com user=vppzvlmbbbwiqc password=45a8e987c8e6dcc38dcd31ba8bd9a56a6f4894513d164f068cab43d071829826 dbname=d2t1tmuqj0iqk2 port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
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
