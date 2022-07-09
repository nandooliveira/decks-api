package conf

import (
	"fmt"
	"log"
	"net/http"
)

type Application struct {
	Port int32
}

func (application *Application) Init() {
	r := InitRoutes()
	addr := fmt.Sprintf(":%v", application.Port)

	log.Printf("Starting server in port %v", application.Port)
	log.Fatal(http.ListenAndServe(addr, r))
}
