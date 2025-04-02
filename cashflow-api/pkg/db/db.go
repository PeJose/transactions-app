package db

import (
	"cashflow/pkg/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	DB, err = gorm.Open(sqlite.Open("ing.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	println("Migrating database...")
	if err := migrate(DB); err != nil {
		panic(err)
	}
}

func migrate(db *gorm.DB) error {
	return db.AutoMigrate(&models.User{}, &models.Company{})
}
