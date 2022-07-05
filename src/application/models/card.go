package models

import "fmt"

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
	Value Rank   `json:"rank"`
	Suit  Suit   `json:"suit"`
	Code  string `json:"code"`
}

func NewCard(rank Rank, suit Suit) Card {
	return Card{rank, suit, fmt.Sprintf("%s%s", rank[0:1], suit[0:1])}
}
