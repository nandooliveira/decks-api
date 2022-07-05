package deck

import "fmt"

type Suit int

type Rank int

const (
	CLUB Suit = iota
	DIAMOND
	HEART
	SPADE
)

const (
	ACE Rank = iota
	TWO
	THREE
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
	TEN
	JACK
	QUEEN
	KING
)

var (
	SUITS = []Suit{CLUB, DIAMOND, HEART, SPADE}
	RANKS = []Rank{ACE, TWO, THREE, FOUR, FIVE, SIX, SEVEN, EIGHT, NINE, TEN, JACK, QUEEN, KING}
)

type Card int

func (c Card) String() string {
	face := ""
	switch c.Face() {
	case 0:
		face = "A"
	case 1:
		face = "2"
	case 2:
		face = "3"
	case 3:
		face = "4"
	case 4:
		face = "5"
	case 5:
		face = "6"
	case 6:
		face = "7"
	case 7:
		face = "8"
	case 8:
		face = "9"
	case 9:
		face = "T"
	case 10:
		face = "J"
	case 11:
		face = "Q"
	case 12:
		face = "K"
	}
	suit := ""
	switch c.Suit() {
	case 0:
		suit = "♣"
	case 1:
		suit = "♦"
	case 2:
		suit = "♥"
	case 3:
		suit = "♠"
	}
	return fmt.Sprintf("%s%s", face, suit)
}

func (c Card) Face() int {
	return int(c / 4)
}

func (c Card) Suit() int {
	return int(c % 4)
}

func NewCard(rank Rank, suit Suit) Card {
	return Card(int(rank)*4 + int(suit))
}
