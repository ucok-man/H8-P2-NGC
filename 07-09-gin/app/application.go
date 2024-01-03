package app

import (
	"log"
	"os"

	"github.com/ucok-man/H8-P2-NGC/07-09-gin/config"
	"github.com/ucok-man/H8-P2-NGC/07-09-gin/internal/repo"
)

type Application struct {
	Config *config.Config
	Log    *log.Logger
	Repo   *repo.Services
}

func New() *Application {
	app := &Application{
		Config: config.New(),
		Log:    log.New(os.Stdout, "APP", log.Ldate|log.Ltime),
	}
	app.Repo = repo.New(app.Config.GetDBConn())

	return app
}
