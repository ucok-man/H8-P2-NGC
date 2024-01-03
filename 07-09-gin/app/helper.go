package app

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ucok-man/H8-P2-NGC/07-09-gin/internal/model"
)

func (app *Application) getCurrentUser(ctx *gin.Context) *model.Store {
	obj, exists := ctx.Get(app.Config.GetUserCtxKey())
	if !exists {
		panic("[app.getCurrentUser]: user should exists!")
	}

	user, ok := obj.(*model.Store)
	if !ok {
		panic("[app.getCurrentUser]: user should be *model.Store")
	}
	return user
}

func (app *Application) getParamId(ctx *gin.Context) (uint, error) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil || id < 1 {
		return uint(0), fmt.Errorf("invalid id parameter")
	}

	return uint(id), nil
}
