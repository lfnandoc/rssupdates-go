package main

import (
	"gorm.io/gorm"

	"gorm.io/driver/sqlite"
)

var DB *gorm.DB

type Post struct {
	gorm.Model
	Guid string
}

func SetupDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(&Post{})
}
