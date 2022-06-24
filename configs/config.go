package configs

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Load() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("book.db"), &gorm.Config{})

	if err != nil {
		panic(("Connection Faild"))
	}

	return db
}
