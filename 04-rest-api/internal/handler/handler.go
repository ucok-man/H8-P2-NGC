package handler

import "github.com/ucok-man/H8-P2-NGC/04-rest-api/internal/config"

type Handler struct {
	CrimeCase CrimeCaseHandler
	Hero      HeroHandler
	Villain   VillainHandler
}

func New(app *config.Application) *Handler {
	return &Handler{
		CrimeCase: CrimeCaseHandler{app: app},
		Hero:      HeroHandler{app: app},
		Villain:   VillainHandler{app: app},
	}
}
