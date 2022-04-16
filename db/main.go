package db

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func InitDB(dbName string) {
	var err error
	db, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	err = db.AutoMigrate(&Recipe{}, &Ingredient{})
	if err != nil {
		log.Fatalln("Cant Auto-migrate")
		return
	}
}
