package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ucok-man/H8-P2-NGC/06-middleware-and-auth/internal/config"
	"github.com/ucok-man/H8-P2-NGC/06-middleware-and-auth/internal/handler"
	"github.com/ucok-man/H8-P2-NGC/06-middleware-and-auth/internal/middleware"
)

func main() {
	app, cleanfn := config.NewApp()
	defer cleanfn()

	handler := handler.New(app)
	mdware := middleware.New(app)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.Config.Port),
		Handler: routes(handler, mdware, app),
	}

	log.Printf("runing server on port: %v\n", app.Config.Port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
