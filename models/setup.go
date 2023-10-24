package models

import {
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
}
var database *gorm.DB
func ConnectDB() {
	db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/mikti"))
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Product{})

	database = db
}