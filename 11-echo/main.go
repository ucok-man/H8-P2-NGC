package main

import (
	"log"

	"github.com/ucok-man/H8-P2-NGC/11-echo/api"
)

func main() {
	app := api.New()
	defer app.Cleanup()

	if err := app.Routes().Serve(); err != nil {
		log.Fatal(err)
	}
}
