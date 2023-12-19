package config

import (
	"log"
	"os"

	"github.com/ucok-man/H8-P2-NGC/06-middleware-and-auth/internal/contract"
	"github.com/ucok-man/H8-P2-NGC/06-middleware-and-auth/internal/entity"
)

type Application struct {
	Config configuration
	JSON   contract.Contract
	Entity entity.Entity
	Log    *log.Logger
}

func NewApp() (*Application, func()) {
	cfg := NewConfig()

	db, err := OpenConn(cfg.DBdsn)
	if err != nil {
		log.Fatal(err)
	}

	app := &Application{
		Config: cfg,
		JSON:   contract.New(),
		Entity: entity.New(db),
		Log:    log.New(os.Stdout, "", 0),
	}

	return app, func() {
		defer db.Close()
	}
}
