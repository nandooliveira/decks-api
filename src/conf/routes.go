package conf

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nandooliveira/deck-api/src/application/controllers"
)

func InitRoutes() *mux.Router {
	r := mux.NewRouter()

	// home routes
	homeController := controllers.HomeController{}
	r.HandleFunc("/", homeController.Index).Methods("GET")

	// decks routes
	decksController := controllers.DecksController{}
	r.HandleFunc("/decks", decksController.Create).Methods("POST")
	r.HandleFunc("/decks/{id}", decksController.Open).Methods("GET")
	r.HandleFunc("/decks/{id}/cards", decksController.DrawCards).Methods("GET")

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
