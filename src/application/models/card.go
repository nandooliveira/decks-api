package models

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Suit string

type Rank string

const (
	CLUBS    = Suit("CLUBS")
	DIAMONDS = Suit("DIAMONDS")
	HEARTS   = Suit("HEARTS")
	SPADES   = Suit("SPADES")
)

const (
	ACE   = Rank("ACE")
	TWO   = Rank("2")
	THREE = Rank("3")
	FOUR  = Rank("4")
	FIVE  = Rank("5")
	SIX   = Rank("6")
	SEVEN = Rank("7")
	EIGHT = Rank("8")
	NINE  = Rank("9")
	TEN   = Rank("10")
	JACK  = Rank("JACK")
	QUEEN = Rank("QUEEN")
	KING  = Rank("KING")
)

var (
	SUITS = []Suit{CLUBS, DIAMONDS, HEARTS, SPADES}
	RANKS = []Rank{ACE, TWO, THREE, FOUR, FIVE, SIX, SEVEN, EIGHT, NINE, TEN, JACK, QUEEN, KING}
)

type Card struct {
	gorm.Model `json:"-"`
	Rank       Rank   `json:"value"`
	Suit       Suit   `json:"suit"`
	Code       string `json:"code"`
	DeckID     uint   `json:"-"`
}

func NewCard(rank Rank, suit Suit) Card {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	card := Card{Rank: rank, Suit: suit, Code: fmt.Sprintf("%s%s", rank[0:1], suit[0:1])}
	db.Create(&card)
	return card
}

func DeleteCard(card Card) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.Delete(&card)
}
