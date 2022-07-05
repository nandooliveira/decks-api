package models

import (
	"math/rand"

	"github.com/google/uuid"
)

type Deck struct {
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

	deck := Deck{uuid.New().String(), cards, opt.Decks, opt.Shuffled}
	if opt.Shuffled {
		deck.Shuffle()
	}

	return &deck, nil
}

func (this *Deck) Shuffle() {
	N := len(this.Cards)
	for i := 0; i < N; i++ {
		r := i + rand.Intn(N-i)
		this.Cards[r], this.Cards[i] = this.Cards[i], this.Cards[r]
	}
}

func Unshuffled(o *Options) {
	o.Shuffled = false
}
