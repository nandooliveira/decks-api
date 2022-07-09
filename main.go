package main

import (
	"github.com/nandooliveira/deck-api/src/application/models"
	"github.com/nandooliveira/deck-api/src/conf"
)

func main() {
	// Migrate the schema
	models.DbManager.DB.AutoMigrate(&models.Card{})
	models.DbManager.DB.AutoMigrate(&models.Deck{})

	app := conf.Application{Port: 8000}
	app.Init()
}
