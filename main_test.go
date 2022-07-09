package main

import (
	"github.com/nandooliveira/deck-api/src/application/models"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	models.DbManager.DB.AutoMigrate(&models.Card{})
	models.DbManager.DB.AutoMigrate(&models.Deck{})

	code := m.Run()
	clearTable()
	os.Exit(code)
}

func clearTable() {
	models.DbManager.DB.Exec("DELETE FROM cards")
	models.DbManager.DB.Exec("DELETE FROM decks")
	models.DbManager.DB.Exec("ALTER SEQUENCE decks_id_seq RESTART WITH 1")
	models.DbManager.DB.Exec("ALTER SEQUENCE cards_id_seq RESTART WITH 1")
}
