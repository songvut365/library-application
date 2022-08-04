package database

import (
	"library-app/model"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	dsn := "host=localhost user=postgres password=1234 dbname=library port=5432 sslmode=disable TimeZone=Asia/Bangkok"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	DB.AutoMigrate(
		&model.Book{},
		&model.BookDetail{},
		&model.Member{},
		&model.Borrow{},
	)
}
