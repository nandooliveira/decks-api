package models

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type Manager struct {
	DB *gorm.DB
}

var DbManager *Manager

func init() {
	dsn := fmt.Sprintf(
		"host=localhost user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=America/Sao_Paulo",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	DbManager = &Manager{DB: db}
}

func ClearTables() {
	DbManager.DB.Exec("DELETE FROM decks")
	DbManager.DB.Exec("DELETE FROM cards")
}
