package main

import (
	"github.com/nandooliveira/deck-api/src/application/models"
	"github.com/nandooliveira/deck-api/src/conf"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	db.AutoMigrate(&models.Card{})
	db.AutoMigrate(&models.Deck{})

	app := conf.Application{Port: 8000}
	app.Init()
}
