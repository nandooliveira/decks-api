package helpers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/nandooliveira/deck-api/src/application/schemas"
)

func RespondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func CheckErrors(w http.ResponseWriter, err schemas.Error) {
	log.Println(err)

	switch err.Type {
	case InternalError:
		RespondWithJson(w, http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"error":   "Internal Error.",
		})
	}
}
