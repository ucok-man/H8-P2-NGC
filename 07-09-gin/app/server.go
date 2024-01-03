package app

import (
	"net/http"
	"time"
)

func (app *Application) Start() error {
	srv := &http.Server{
		Addr:         app.Config.Addr,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	return srv.ListenAndServe()
}
