package api

import (
	"fmt"
)

func (app *Application) Serve() error {
	err := app.echo.Start(fmt.Sprintf(":%d", app.config.ApiPort))
	if err != nil {
		return err
	}
	return nil
}
