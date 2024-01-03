package contract

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendSuccess(ctx *gin.Context, code int, message string, data any) {
	ctx.JSON(code, APIResponse{
		Success: true,
		Message: message,
		Data:    data,
		Error:   nil,
	})
}

func StatusOK(ctx *gin.Context, message string, data any) {
	sendSuccess(ctx, http.StatusOK, message, data)
}

func StatusCreated(ctx *gin.Context, message string, data any) {
	sendSuccess(ctx, http.StatusCreated, message, data)
}
