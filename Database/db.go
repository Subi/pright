package database

import (
	"log"
	model "pright/Model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dsn string = "postgres://postgres:mynewpassword@localhost:5432"

func NewDb() *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database : %s", err)
	}
	db.AutoMigrate(&model.Hotel{})

	return db
}
