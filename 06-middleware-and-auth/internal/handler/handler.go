package handler

import "github.com/ucok-man/H8-P2-NGC/06-middleware-and-auth/internal/config"

type Handler struct {
	User   UserHandler
	Recipe RecipeHandler
}

func New(app *config.Application) *Handler {
	return &Handler{
		User:   UserHandler{app},
		Recipe: RecipeHandler{app},
	}
}
