package config

import (
	"log"

	"github.com/ucok-man/H8-P2-NGC/03-router/entity"
)

type Application struct {
	Addr   string
	Entity *entity.Entity
}

func NewApp() (*Application, func()) {
	db, err := OpenConn()
	if err != nil {
		log.Fatal(err)
	}

	app := &Application{
		Addr:   "localhost:8000",
		Entity: entity.New(db),
	}

	return app, func() {
		defer db.Close()
	}
}
