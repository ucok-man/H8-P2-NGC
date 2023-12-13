package main

import (
	"log"
	"net/http"

	"github.com/ucok-man/H8-P2-NGC/04-rest-api/internal/config"
	"github.com/ucok-man/H8-P2-NGC/04-rest-api/internal/handler"
)

func main() {
	app, cleanfn := config.NewApp()
	defer cleanfn()

	handler := handler.New(app)

	srv := &http.Server{
		Addr:    app.Addr,
		Handler: routes(handler, app),
	}

	log.Printf("runing server on %v...\n", app.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
