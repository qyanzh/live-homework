package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"live/internal/entity"
	"log"
)

var DbClient *gorm.DB

func init() {
	var err error
	DbClient, err = gorm.Open(mysql.Open("qyanzh:qyanzh@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}

	DbClient.AutoMigrate(
		&entity.User{},
		&entity.Room{},
		&entity.Good{},
		&entity.Tx{},
	)
}
