package conf

import (
	"fmt"
	"log"
	"net/http"
)

type Application struct {
	Port int32
}

func (this *Application) Init() {
	r := InitRoutes()
	addr := fmt.Sprintf(":%v", this.Port)

	log.Printf("Starting server in port %v", this.Port)
	log.Fatal(http.ListenAndServe(addr, r))
}
