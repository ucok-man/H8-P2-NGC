package main

import (
	"log"
	"net/http"

	"github.com/ucok-man/H8-P2-NGC/02-web-server/config"
	"github.com/ucok-man/H8-P2-NGC/02-web-server/entity"
)

type application struct {
	addr   string
	entity *entity.Entity
}

func main() {
	db, err := config.OpenConn()
	if err != nil {
		log.Fatal(err)
	}

	app := &application{
		addr:   "localhost:8000",
		entity: entity.New(db),
	}

	mux := http.NewServeMux()

	fileserver := http.FileServer(http.Dir("./02-web-server/assets/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileserver))

	mux.HandleFunc("/heroes", app.GetHeroes)
	mux.HandleFunc("/villains", app.GetVillains)

	srv := &http.Server{
		Addr:    app.addr,
		Handler: mux,
	}

	log.Printf("runing server on %v...\n", app.addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
