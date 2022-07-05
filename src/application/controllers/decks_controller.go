package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/nandooliveira/deck-api/src/application/helpers"
	"github.com/nandooliveira/deck-api/src/application/models"
)

type DecksController struct{}

func (this *DecksController) Create(w http.ResponseWriter, r *http.Request) {
	shuffled := r.URL.Query().Get("shuffled")
	deck, err := models.NewDeck(func(o *models.Options) {
		o.Shuffled, _ = strconv.ParseBool(shuffled)
	})

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
