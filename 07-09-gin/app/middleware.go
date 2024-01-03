package app

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ucok-man/H8-P2-NGC/07-09-gin/contract"
	"github.com/ucok-man/H8-P2-NGC/07-09-gin/internal/jwt"
	"github.com/ucok-man/H8-P2-NGC/07-09-gin/internal/repo"
)

func (app *Application) WithErrorHandle(ctx *gin.Context) {
	ctx.Next()

	if len(ctx.Errors) > 0 {
		for _, err := range ctx.Errors {
			app.Log.Printf("[ERROR]: %v\n", err)
		}
	}
}

func (app *Application) WithLogin(ctx *gin.Context) {
	authorizationHeader := ctx.GetHeader("Authorization")
	if authorizationHeader == "" {
		contract.ErrInvalidAuthenticationToken(ctx)
		return
	}

	headerParts := strings.Split(authorizationHeader, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		contract.ErrInvalidAuthenticationToken(ctx)
		return
	}
	tokenstr := headerParts[1]

	var claim jwt.JWTClaim
	err := jwt.DecodeToken(tokenstr, &claim, app.Config.GetJwtSecret())
	if err != nil {
		contract.ErrInvalidAuthenticationToken(ctx)
		return
	}

	user, err := app.Repo.User.GetByID(claim.UserID)
	if err != nil {
		switch {
		case errors.Is(err, repo.ErrRecordNotFound):
			contract.ErrInvalidCredentials(ctx)
		default:
			contract.ErrInternalServer(ctx, err)
		}
		return
	}

	// set context
	ctx.Set(app.Config.GetUserCtxKey(), user)
	ctx.Next()
}
