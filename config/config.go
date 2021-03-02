package config

import (
	"Final/structs"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	// "github.com/jinzhu/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:@tcp(127.0.0.1:3306)/assigment3?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	db.AutoMigrate(structs.Label{}, structs.Project{}, structs.Status{}, structs.Task{})

	fmt.Println("Connected to Database")

	DB = db
}
