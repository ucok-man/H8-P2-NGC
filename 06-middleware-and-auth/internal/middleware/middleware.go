package middleware

import "github.com/ucok-man/H8-P2-NGC/06-middleware-and-auth/internal/config"

type Middleware struct {
	*config.Application
}

func New(app *config.Application) *Middleware {
	return &Middleware{app}
}
