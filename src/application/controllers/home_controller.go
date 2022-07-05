package controllers

import (
	"net/http"

	"github.com/nandooliveira/deck-api/src/application/helpers"
)

type HomeController struct{}

func (this *HomeController) Index(w http.ResponseWriter, r *http.Request) {
	helpers.RespondWithJson(w, http.StatusOK, map[string]interface{}{"message": "OK"})
}
