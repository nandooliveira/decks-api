package models

import (
	"math/rand"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Deck struct {
	gorm.Model
	Uuid          string `json:"uuid"`
	Cards         []Card `json:"cards"`
	NumberOfDecks int    `json:"number_of_decks"`
	Shuffled      bool   `json:"shuffled"`
}

type Options struct {
	Shuffled bool
	Cards    []Card
	Ranks    []Rank
	Suits    []Suit
	Decks    int
}

func NewDeck(options ...func(*Options)) (*Deck, error) {
	opt := Options{Shuffled: true, Ranks: RANKS, Suits: SUITS, Decks: 1, Cards: []Card{}}
	for _, option := range options {
		option(&opt)
	}

	cards := opt.Cards

	if len(cards) == 0 {
		cards = make([]Card, len(opt.Suits)*len(opt.Ranks)*opt.Decks)
		index := 0
		for i := 0; i < opt.Decks; i++ {
			for _, suit := range opt.Suits {
				for _, rank := range opt.Ranks {
					cards[index] = NewCard(rank, suit)
					index++
				}
			}
		}
	}

	deck := Deck{Uuid: uuid.New().String(), Cards: cards, NumberOfDecks: opt.Decks, Shuffled: opt.Shuffled}
	DbManager.DB.Create(&deck)

	if opt.Shuffled {
		deck.Shuffle()
	}

	return &deck, nil
}

func FindDeck(uuid string) Deck {
	var deck Deck
	DbManager.DB.Preload("Cards").First(&deck, "uuid = ?", uuid)

	return deck
}

func DrawCards(uuid string, count int) []Card {
	deck := FindDeck(uuid)
	cards := deck.Cards[:count]

	for _, card := range cards {
		DeleteCard(card)
	}

	deck.Cards = deck.Cards[count:]

	return cards
}

func (d *Deck) Shuffle() {
	N := len(d.Cards)
	for i := 0; i < N; i++ {
		r := i + rand.Intn(N-i)
		d.Cards[r], d.Cards[i] = d.Cards[i], d.Cards[r]
	}
}

func Unshuffled(o *Options) {
	o.Shuffled = false
}
