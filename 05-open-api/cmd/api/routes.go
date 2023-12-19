package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/ucok-man/H8-P2-NGC/05-open-api/internal/config"
	"github.com/ucok-man/H8-P2-NGC/05-open-api/internal/handler"
)

func routes(handler *handler.Handler, app *config.Application) http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.Error.NotFound)
	router.MethodNotAllowed = http.HandlerFunc(app.Error.MethodNotAllowed)

	// hero
	router.GET("/heroes", handler.Hero.GetAll)
	router.POST("/heroes", handler.Hero.Create)

	// villain
	router.GET("/villains", handler.Villain.GetAll)
	router.POST("/villains", handler.Villain.Create)

	// crimecase
	router.GET("/crimecases", handler.CrimeCase.GetAll)
	router.GET("/crimecases/:id", handler.CrimeCase.GetByID)
	router.POST("/crimecases", handler.CrimeCase.Create)
	router.PUT("/crimecases/:id", handler.CrimeCase.Update)
	router.DELETE("/crimecases/:id", handler.CrimeCase.DeleteByID)

	return router
}
