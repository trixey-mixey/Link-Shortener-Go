package main

import (
	"go/projcet-Adv/internal/link"
	"go/projcet-Adv/internal/stat"
	"go/projcet-Adv/internal/users"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&link.Link{}, &users.User{}, &stat.Stat{})
}
