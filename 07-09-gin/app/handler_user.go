package app

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/ucok-man/H8-P2-NGC/07-09-gin/contract"
	"github.com/ucok-man/H8-P2-NGC/07-09-gin/internal/jwt"
	"github.com/ucok-man/H8-P2-NGC/07-09-gin/internal/repo"
)

func (app *Application) loginHandler(ctx *gin.Context) {
	var input contract.ReqLogin

	if err := ctx.ShouldBindJSON(&input); err != nil {
		contract.ErrBadRequest(ctx, err)
		return
	}

	if err := input.Validate(); err != nil {
		contract.ErrFailedValidation(ctx, err)
		return
	}

	user, err := app.Repo.User.GetByEmail(input.Email)
	if err != nil {
		switch {
		case errors.Is(err, repo.ErrRecordNotFound):
			contract.ErrInvalidCredentials(ctx)
		default:
			contract.ErrInternalServer(ctx, err)
		}
		return
	}

	if err := user.Password.Matches(input.Password); err != nil {
		contract.ErrInvalidCredentials(ctx)
		return
	}

	claims := jwt.NewJWTClaim(user.StoreID)
	token, err := jwt.GenerateToken(&claims, app.Config.GetJwtSecret())
	if err != nil {
		contract.ErrInternalServer(ctx, err)
		return
	}

	contract.StatusOK(ctx, "OK", gin.H{
		"access_token": token,
		"token_type":   "Bearer",
	})
}

func (app *Application) registerHandler(ctx *gin.Context) {
	var input contract.ReqRegister

	if err := ctx.ShouldBindJSON(&input); err != nil {
		contract.ErrBadRequest(ctx, err)
		return
	}

	if err := input.Validate(); err != nil {
		contract.ErrFailedValidation(ctx, err)
		return
	}

	user := input.ToUser()
	if err := app.Repo.User.Insert(user); err != nil {
		switch {
		case errors.Is(err, repo.ErrDuplicateRecord):
			contract.ErrDuplicateRecord(ctx, "email", "already exists")
		default:
			contract.ErrInternalServer(ctx, err)
		}
		return
	}

	claims := jwt.NewJWTClaim(user.StoreID)
	token, err := jwt.GenerateToken(&claims, app.Config.GetJwtSecret())
	if err != nil {
		contract.ErrInternalServer(ctx, err)
		return
	}

	contract.StatusCreated(ctx, "Created", gin.H{
		"access_token": token,
		"token_type":   "Bearer",
		"user":         user,
	})
}
