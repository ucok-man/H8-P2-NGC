package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
	"github.com/ucok-man/H8-P2-NGC/11-echo/internal/contract"
	"github.com/ucok-man/H8-P2-NGC/11-echo/internal/jwt"
	"github.com/ucok-man/H8-P2-NGC/11-echo/internal/pkg/validator"
	"github.com/ucok-man/H8-P2-NGC/11-echo/internal/repo"
)

func (app *Application) registerHandler(ctx echo.Context) error {
	var input contract.ReqRegister
	if err := ctx.Bind(&input); err != nil {
		return app.ErrBadRequest(ctx, err)
	}

	if err := ctx.Validate(&input); err != nil {
		return app.ErrFailedValidation(ctx, err)
	}

	user, err := input.ToUser()
	if err != nil {
		return app.ErrInternalServer(ctx, err)
	}

	if err = app.repo.User.Insert(user); err != nil {
		switch {
		case errors.Is(err, repo.ErrorDuplicateRecord):
			return app.ErrFailedValidation(ctx, validator.AddError("email", "already exists"))
		default:
			return app.ErrInternalServer(ctx, err)
		}
	}

	claims := jwt.NewJWTClaim(user.UserID)
	token, err := jwt.GenerateToken(&claims, app.config.JwtSecret)
	if err != nil {
		app.ErrInternalServer(ctx, err)
	}

	return ctx.JSON(http.StatusCreated, gin.H{
		"access_token": token,
		"message":      "success register",
		"data":         user,
	})

}

func (app *Application) loginHandler(ctx echo.Context) error {
	var input contract.ReqLogin
	if err := ctx.Bind(&input); err != nil {
		return app.ErrBadRequest(ctx, err)
	}

	if err := ctx.Validate(&input); err != nil {
		return app.ErrFailedValidation(ctx, err)
	}

	user, err := app.repo.User.GetByUsername(input.Username)
	if err != nil {
		switch {
		case errors.Is(err, repo.ErrorRecordNotFound):
			return app.ErrInvalidCredentials(ctx)
		default:
			return app.ErrInternalServer(ctx, err)
		}
	}

	if err := user.MatchesPassword(input.Password); err != nil {
		return app.ErrInvalidCredentials(ctx)
	}

	claims := jwt.NewJWTClaim(user.UserID)
	token, err := jwt.GenerateToken(&claims, app.config.JwtSecret)
	if err != nil {
		return app.ErrInternalServer(ctx, err)
	}

	return ctx.JSON(http.StatusOK, gin.H{
		"access_token": token,
		"message":      "success login",
	})
}
