package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/nandooliveira/deck-api/src/application/helpers"
	"github.com/nandooliveira/deck-api/src/application/models"
)

var (
	translateSuits = map[string]models.Suit{
		"C": models.CLUBS,
		"D": models.DIAMONDS,
		"H": models.HEARTS,
		"S": models.SPADES,
	}

	translateRanks = map[string]models.Rank{
		"A":  models.ACE,
		"2":  models.TWO,
		"3":  models.THREE,
		"4":  models.FOUR,
		"5":  models.FIVE,
		"6":  models.SIX,
		"7":  models.SEVEN,
		"8":  models.EIGHT,
		"9":  models.NINE,
		"10": models.TEN,
		"J":  models.JACK,
		"Q":  models.QUEEN,
		"K":  models.KING,
	}
)

type DecksController struct{}

func (this *DecksController) Create(w http.ResponseWriter, r *http.Request) {
	shuffled := r.URL.Query().Get("shuffled")

	codeCards := r.URL.Query().Get("cards")

	options := func(o *models.Options) {
		o.Shuffled, _ = strconv.ParseBool(shuffled)
		if codeCards != "" {
			splittedCards := strings.Split(codeCards, ",")

			for _, cardCode := range splittedCards {
				o.Cards = append(o.Cards, models.NewCard(translateRanks[cardCode[0:1]], translateSuits[cardCode[1:2]]))
			}
		}
	}

	deck, err := models.NewDeck(options)

	if err != nil {
		handleParamError(w, err.Error())
		return
	}

	helpers.RespondWithJson(w, http.StatusOK, map[string]interface{}{
		"deck_id":   deck.Uuid,
		"shuffled":  deck.Shuffled,
		"remaining": len(deck.Cards),
	})
}

func (this *DecksController) Open(w http.ResponseWriter, r *http.Request) {
	paramId := mux.Vars(r)["id"]
	id, errParam := strconv.ParseInt(paramId, 10, 64)

	if errParam != nil {
		handleParamError(w, "Invalid Deck ID")
		return
	}

	helpers.RespondWithJson(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"message": fmt.Sprintf("[%s] Open Deck: %d", uuid.New(), id),
	})
}

func (this *DecksController) DrawCards(w http.ResponseWriter, r *http.Request) {
	paramId := mux.Vars(r)["id"]
	id, errParam := strconv.ParseInt(paramId, 10, 64)

	if errParam != nil {
		handleParamError(w, "Invalid Deck ID")
		return
	}

	helpers.RespondWithJson(w, http.StatusOK, map[string]interface{}{
		"success": true,
		"message": fmt.Sprintf("[%s] Draw Deck Cards: %d", uuid.New(), id),
	})
}

func handleParamError(w http.ResponseWriter, err string) {
	helpers.RespondWithJson(w, http.StatusBadRequest, map[string]interface{}{
		"success": false,
		"error":   err,
	})
}
