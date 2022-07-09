package conf

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nandooliveira/deck-api/src/application/controllers"
)

func InitRoutes() *mux.Router {
	r := mux.NewRouter()

	// decks routes
	decksController := controllers.DecksController{}
	r.HandleFunc("/decks", decksController.Create).Methods("POST")
	r.HandleFunc("/decks/{uuid}", decksController.Open).Methods("GET")
	r.HandleFunc("/decks/{uuid}/{count}", decksController.DrawCards).Methods("PUT")

	r.Use(logMw)
	http.Handle("/", r)

	return r
}

func logMw(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Printf("%v REQUESTING: %v", r.RemoteAddr, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
