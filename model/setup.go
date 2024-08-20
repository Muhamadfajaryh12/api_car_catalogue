package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase (){
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/car_catalogue"))
	if err != nil{
		panic(err)
	}

	database.AutoMigrate(&Product{},&Category{},&User{})
	DB=database
}

