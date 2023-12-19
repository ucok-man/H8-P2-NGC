package config

import (
	"log"

	"github.com/ucok-man/H8-P2-NGC/05-open-api/internal/entity"
	"github.com/ucok-man/H8-P2-NGC/05-open-api/internal/error_response"
	"github.com/ucok-man/H8-P2-NGC/05-open-api/internal/util"
)

type Application struct {
	Addr   string
	Util   *util.Utility
	Error  error_response.Error
	Entity *entity.Entity
}

func NewApp() (*Application, func()) {
	db, err := OpenConn()
	if err != nil {
		log.Fatal(err)
	}

	util := util.New()
	cfg := &Application{
		Addr:   "localhost:8000",
		Entity: entity.New(db),
		Error:  *error_response.New(util),
		Util:   util,
	}

	return cfg, func() {
		defer db.Close()
	}
}
