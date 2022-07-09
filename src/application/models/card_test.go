package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCard(t *testing.T) {
	card := NewCard(ACE, CLUBS)
	assert.Equal(t, Rank("ACE"), card.Rank)

	ClearTables()
}
