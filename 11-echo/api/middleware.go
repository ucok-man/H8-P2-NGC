package api

import (
	"errors"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/ucok-man/H8-P2-NGC/11-echo/internal/jwt"
	"github.com/ucok-man/H8-P2-NGC/11-echo/internal/repo"
)

func (app *Application) WithLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		authorizationHeader := ctx.Request().Header.Get("Authorization")
		if authorizationHeader == "" {
			return app.ErrInvalidCredentials(ctx)
		}

		headerParts := strings.Split(authorizationHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			return app.ErrInvalidCredentials(ctx)
		}
		tokenstr := headerParts[1]

		var claim jwt.JWTClaim
		err := jwt.DecodeToken(tokenstr, &claim, app.config.JwtSecret)
		if err != nil {
			return app.ErrInvalidCredentials(ctx)
		}

		user, err := app.repo.User.GetByID(claim.UserID)
		if err != nil {
			switch {
			case errors.Is(err, repo.ErrorRecordNotFound):
				return app.ErrInvalidCredentials(ctx)
			default:
				return app.ErrInternalServer(ctx, err)
			}
		}

		// set context
		ctx.Set(app.ctxkey.user, user)
		return next(ctx)
	}
}
