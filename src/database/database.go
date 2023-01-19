package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDb(dbAddress string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dbAddress), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}

	seedDB(db)
	return db
}
