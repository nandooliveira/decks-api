package deck

import (
	"math/rand"
)

type Deck struct {
	Cards         []Card
	NumberOfDecks int
}

type Options struct {
	Shuffled bool
	Cards    []Card
	Ranks    []Rank
	Suits    []Suit
	Decks    int
}

func New(options ...func(*Options)) (*Deck, error) {
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

	deck := Deck{cards, opt.Decks}
	if opt.Shuffled {
		deck.Shuffle()
	}

	return &deck, nil
}

func (d *Deck) Shuffle() {
	N := len(d.Cards)
	for i := 0; i < N; i++ {
		r := i + rand.Intn(N-i)
		d.Cards[r], d.Cards[i] = d.Cards[i], d.Cards[r]
	}
}
