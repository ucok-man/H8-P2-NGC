package contract

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendErr(ctx *gin.Context, code int, msg string, details any) {
	ctx.Abort()
	ctx.JSON(code, APIResponse{
		Success: false,
		Message: msg,
		Data:    nil,
		Error: &ApiError{
			StatusCode: http.StatusText(code),
			Details:    details,
		},
	})
}

func ErrInternalServer(ctx *gin.Context, err error) {
	ctx.Error(err)
	msg := "the server encountered a problem and could not process your request"
	sendErr(ctx, http.StatusInternalServerError, msg, nil)
}

func ErrNotFound(ctx *gin.Context) {
	msg := "the requested resource could not be found"
	sendErr(ctx, http.StatusNotFound, msg, nil)
}

func ErrMethodNotAllowed(ctx *gin.Context) {
	msg := fmt.Sprintf("the %s method is not supported for this resource", ctx.Request.Method)
	sendErr(ctx, http.StatusMethodNotAllowed, msg, nil)
}

func ErrBadRequest(ctx *gin.Context, err error) {
	sendErr(ctx, http.StatusBadRequest, err.Error(), nil)
}

func ErrFailedValidation(ctx *gin.Context, errors map[string]string) {
	msg := "unable to process entity because some malformed value"
	sendErr(ctx, http.StatusUnprocessableEntity, msg, errors)
}

func ErrDuplicateRecord(ctx *gin.Context, keydetails, msgdetails string) {
	msg := "unable to process entity because some malformed value"
	sendErr(ctx, http.StatusUnprocessableEntity, msg, map[string]string{keydetails: msgdetails})
}

func ErrInvalidCredentials(ctx *gin.Context) {
	msg := "invalid authentication credentials"
	sendErr(ctx, http.StatusUnauthorized, msg, nil)
}

func ErrInvalidAuthenticationToken(ctx *gin.Context) {
	ctx.Header("WWW-Authenticate", "Bearer")

	msg := "invalid or missing authentication token"
	sendErr(ctx, http.StatusUnauthorized, msg, nil)
}

func ErrAuthenticationRequired(ctx *gin.Context) {
	msg := "you must be authenticated to access this resource"
	sendErr(ctx, http.StatusUnauthorized, msg, nil)
}

func ErrNotPermittedResource(ctx *gin.Context) {
	msg := "your user account doesn't have the necessary permissions to access this resource"
	sendErr(ctx, http.StatusForbidden, msg, nil)
}
