package db

import (
	"go/projcet-Adv/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	*gorm.DB
}

func NewDb(conf *configs.Config) *Db {
	db, err := gorm.Open(postgres.Open(conf.Db.Dsn))
	if err != nil {
		panic(err)
	}
	return &Db{db}
}
