package database

import (
	"library-app/model"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	DSN := os.Getenv("POSTGRES_DSN")

	var err error
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
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
