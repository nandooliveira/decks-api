package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/nandooliveira/deck-api/src/application/models"
	"github.com/steinfletcher/apitest"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateWithoutParams(t *testing.T) {
	r := mux.NewRouter()

	decksController := DecksController{}
	r.HandleFunc("/decks", decksController.Create).Methods("POST")
	ts := httptest.NewServer(r)
	defer ts.Close()
	t.Run("found", func(t *testing.T) {
		apitest.New().
			Handler(r).
			Post("/decks").
			Expect(t).
			Assert(jsonpath.Equal("$.remaining", 52.0)).
			Assert(jsonpath.Equal("$.shuffled", false)).
			Status(http.StatusOK).
			End()
	})
	models.ClearTables()
}

func TestCreateWithCards(t *testing.T) {
	r := mux.NewRouter()

	decksController := DecksController{}
	r.HandleFunc("/decks", decksController.Create).Methods("POST")
	ts := httptest.NewServer(r)
	defer ts.Close()
	t.Run("found", func(t *testing.T) {
		apitest.New().
			Handler(r).
			Intercept(func(req *http.Request) {
				req.URL.RawQuery = "cards=AS,KD,AC,2C,KH"
			}).
			Post("/decks").
			Expect(t).
			Assert(jsonpath.Equal("$.remaining", 5.0)).
			Assert(jsonpath.Equal("$.shuffled", false)).
			Status(http.StatusOK).
			End()
	})
	models.ClearTables()
}

func TestCreateWithShuffled(t *testing.T) {
	r := mux.NewRouter()

	decksController := DecksController{}
	r.HandleFunc("/decks", decksController.Create).Methods("POST")
	ts := httptest.NewServer(r)
	defer ts.Close()
	t.Run("found", func(t *testing.T) {
		apitest.New().
			Handler(r).
			Intercept(func(req *http.Request) {
				req.URL.RawQuery = "shuffled=true"
			}).
			Post("/decks").
			Expect(t).
			Assert(jsonpath.Equal("$.remaining", 52.0)).
			Assert(jsonpath.Equal("$.shuffled", true)).
			Status(http.StatusOK).
			End()
	})
	models.ClearTables()
}

func TestOpen(t *testing.T) {
	r := mux.NewRouter()

	deck, _ := models.NewDeck()

	decksController := DecksController{}
	r.HandleFunc("/decks/{uuid}", decksController.Open).Methods("GET")
	ts := httptest.NewServer(r)
	defer ts.Close()
	t.Run("found", func(t *testing.T) {
		apitest.New().
			Handler(r).
			Get(fmt.Sprintf("/decks/%s", deck.Uuid)).
			Expect(t).
			Assert(jsonpath.Equal("$.deck_id", deck.Uuid)).
			Assert(jsonpath.Equal("$.remaining", 52.0)).
			Assert(jsonpath.Equal("$.shuffled", true)).
			Status(http.StatusOK).
			End()
	})
	models.ClearTables()
}

func TestDrawCards(t *testing.T) {
	r := mux.NewRouter()

	deck, _ := models.NewDeck()

	decksController := DecksController{}
	r.HandleFunc("/decks/{uuid}/{count}", decksController.DrawCards).Methods("PUT")
	ts := httptest.NewServer(r)
	defer ts.Close()
	t.Run("found", func(t *testing.T) {
		apitest.New().
			Handler(r).
			Put(fmt.Sprintf("/decks/%s/%d", deck.Uuid, 3)).
			Expect(t).
			Status(http.StatusOK).
			End()
	})

	reloadedDeck := models.FindDeck(deck.Uuid)
	assert.Equal(t, 49, len(reloadedDeck.Cards))
	models.ClearTables()
}
