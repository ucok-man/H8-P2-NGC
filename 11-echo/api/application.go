package api

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/ucok-man/H8-P2-NGC/11-echo/internal/config"
	"github.com/ucok-man/H8-P2-NGC/11-echo/internal/pkg/json"
	"github.com/ucok-man/H8-P2-NGC/11-echo/internal/repo"
)

type Application struct {
	echo   *echo.Echo
	config *config.Config
	repo   *repo.RepoServices
	ctxkey struct {
		user string
	}
}

func New() *Application {
	app := &Application{
		echo: echo.New(),
		ctxkey: struct{ user string }{
			user: "user",
		},
	}

	cfg, err := config.New()
	if err != nil {
		if err, ok := err.(config.ErrorMissingValue); ok {
			config.ShowHelp(err.Listenv, err.Errors)
		}
		app.echo.Logger.Fatal(err, "error initializing config", nil)
	}
	app.config = cfg
	app.repo = repo.New(app.config.GetDBConn())

	// setup echo
	app.echo.Logger.SetOutput(os.Stdout)
	app.echo.Logger.SetLevel(log.DEBUG)
	app.echo.Logger.SetPrefix("APP")
	app.echo.Validator = NewAppValidator()
	app.echo.JSONSerializer = json.JSONSerializer{}
	app.echo.HTTPErrorHandler = app.httpErrorHandler

	return app
}

func (app *Application) Cleanup() {
	db, err := app.config.GetDBConn().DB()
	if err != nil {
		app.echo.Logger.Fatal("[api.Cleanup] error getting underlying DB")
	}
	db.Close()
}
