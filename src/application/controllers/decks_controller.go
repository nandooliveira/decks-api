package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nandooliveira/deck-api/src/application/helpers"
)

type DecksController struct{}

func (this *DecksController) Create(w http.ResponseWriter, r *http.Request) {
	helpers.RespondWithJson(w, http.StatusOK, map[string]interface{}{"message": "Create Deck"})
}

func (this *DecksController) Open(w http.ResponseWriter, r *http.Request) {
	paramId := mux.Vars(r)["id"]
	id, errParam := strconv.ParseInt(paramId, 10, 64)

	if errParam != nil {
		handleParamError(w, "Invalid Deck ID")
		return
	}

	helpers.RespondWithJson(w, http.StatusOK, map[string]interface{}{"message": fmt.Sprintf("Open Deck: %d", id)})
}

func (this *DecksController) DrawCards(w http.ResponseWriter, r *http.Request) {
	paramId := mux.Vars(r)["id"]
	id, errParam := strconv.ParseInt(paramId, 10, 64)

	if errParam != nil {
		handleParamError(w, "Invalid Deck ID")
		return
	}

	helpers.RespondWithJson(w, http.StatusOK, map[string]interface{}{"message": fmt.Sprintf("Draw Deck Cards: %d", id)})
}

func handleParamError(w http.ResponseWriter, err string) {
	helpers.RespondWithJson(w, http.StatusBadRequest, map[string]interface{}{
		"success": false,
		"error":   err,
	})
}
