package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	database, err := gorm.Open(sqlite.Open("server.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to Database")
	}

	// err = database.AutoMigrate(&Book{})
	// if err != nil {
	// 	return
	// }

	DB = database
}
