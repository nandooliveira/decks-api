package conf

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Application struct {
	Port   int32
	Router *mux.Router
}

func (application *Application) Init() {
	application.Router = InitRoutes()
	addr := fmt.Sprintf(":%v", application.Port)

	log.Printf("Starting server in port %v", application.Port)
	log.Fatal(http.ListenAndServe(addr, application.Router))
}
