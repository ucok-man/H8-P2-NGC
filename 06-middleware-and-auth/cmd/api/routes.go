package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"github.com/ucok-man/H8-P2-NGC/06-middleware-and-auth/internal/config"
	"github.com/ucok-man/H8-P2-NGC/06-middleware-and-auth/internal/handler"
	"github.com/ucok-man/H8-P2-NGC/06-middleware-and-auth/internal/middleware"
)

func routes(handler *handler.Handler, md *middleware.Middleware, app *config.Application) http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.JSON.Error.NotFound)
	router.MethodNotAllowed = http.HandlerFunc(app.JSON.Error.MethodNotAllowed)

	// middleware
	globalMD := alice.New(md.Logging)
	authenticationMD := alice.New(md.RequiredLogin)
	authorizationMD := alice.New(md.RequiredLogin, md.RequiredSuperUser)

	// user
	router.HandlerFunc(http.MethodPost, "/login", handler.User.Login)
	router.HandlerFunc(http.MethodPost, "/register", handler.User.Register)

	// recipe
	router.Handler(http.MethodGet, "/recipes", authenticationMD.ThenFunc(handler.Recipe.GetAll))
	router.Handler(http.MethodGet, "/recipes/:id", authenticationMD.ThenFunc(handler.Recipe.GetByID))
	router.Handler(http.MethodPut, "/recipes/:id", authenticationMD.ThenFunc(handler.Recipe.UpdateByID))

	router.Handler(http.MethodPost, "/recipes", authorizationMD.ThenFunc(handler.Recipe.Create))
	router.Handler(http.MethodDelete, "/recipes/:id", authorizationMD.ThenFunc(handler.Recipe.DeleteByID))

	return globalMD.Then(router)
}
