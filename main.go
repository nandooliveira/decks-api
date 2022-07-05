package main

import "github.com/nandooliveira/deck-api/src/conf"

func main() {
	app := conf.Application{Port: 8000}
	app.Init()
}
