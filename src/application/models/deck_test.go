package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWithCards(t *testing.T) {
	WithCards := func(o *Options) {
		o.Cards = []Card{
			NewCard(ACE, CLUBS),
			NewCard(ACE, DIAMONDS),
			NewCard(ACE, HEARTS),
			NewCard(ACE, SPADES),
		}
	}
	deck, err := NewDeck(WithCards)

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, 4, len(deck.Cards))

	ClearTables()
}

func TestShuffle(t *testing.T) {
	unshuffled, _ := NewDeck(Unshuffled)
	deck, _ := NewDeck()
	deck.Shuffle()

	foundFalse := false
	for i, card := range unshuffled.Cards {
		if card.Rank != deck.Cards[i].Rank || card.Suit != deck.Cards[i].Suit {
			foundFalse = true
			break
		}
	}

	assert.Equal(t, true, foundFalse)

	ClearTables()
}

func TestEmptyShoe(t *testing.T) {
	Decks := func(o *Options) {
		o.Decks = 0
	}
	shoe, _ := NewDeck(Decks)
	result := shoe.NumberOfDecks
	assert.Equal(t, 0, result, "These should be equal")

	ClearTables()
}

func TestFindDeck(t *testing.T) {
	deck, _ := NewDeck()
	result := FindDeck(deck.Uuid)
	assert.Equal(t, deck.Uuid, result.Uuid, "These should be equal")

	ClearTables()
}

func TestDrawCards(t *testing.T) {
	deck, _ := NewDeck()
	cards := DrawCards(deck.Uuid, 2)
	assert.Equal(t, 2, len(cards), "These should be equal")

	reloaded_deck := FindDeck(deck.Uuid)
	assert.Equal(t, 50, len(reloaded_deck.Cards), "These should be equal")

	ClearTables()
}
