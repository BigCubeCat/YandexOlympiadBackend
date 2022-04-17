package db

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB(dbName string) {
	var err error
	DB, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	err = DB.AutoMigrate(&Recipe{}, &Ingredient{})
	if err != nil {
		log.Fatalln("Cant Auto-migrate")
		return
	}
}
