package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"reviews/biz/model/review"
)

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	//dsn := "gorm:gorm@tcp(10.253.17.27:3306)/review?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "gorm:gorm@tcp(localhost:3306)/review?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&review.Review{})
}
