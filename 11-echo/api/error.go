package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/ucok-man/H8-P2-NGC/11-echo/internal/contract"
)

func (app *Application) logError(ctx echo.Context, err error) {
	app.echo.Logger.Errorj(log.JSON{
		"request_method": ctx.Request().Method,
		"request_url":    ctx.Request().URL.String(),
		"error_message":  err,
	})
}

func (app *Application) httpErrorHandler(err error, ctx echo.Context) {
	if ctx.Response().Committed {
		return
	}

	httperr, ok := err.(*echo.HTTPError)
	if !ok {
		app.logError(ctx, err)
		app.echo.DefaultHTTPErrorHandler(err, ctx)
	}

	if httperr.Code == http.StatusNotFound {
		httperr = app.ErrNotFound(ctx)
	}

	if httperr.Code == http.StatusMethodNotAllowed {
		httperr = app.ErrMethodNotAllowed(ctx)
	}

	var response = contract.ErrorResponse{
		StatusText: http.StatusText(httperr.Code),
		Message:    httperr.Message,
	}

	if httperr.Internal != nil {
		response.Details = httperr.Internal
	}

	err = ctx.JSON(httperr.Code, map[string]any{"error": response})
	if err != nil {
		app.logError(ctx, err)
		ctx.Response().WriteHeader(http.StatusInternalServerError)
	}
}

func (app *Application) ErrNotFound(ctx echo.Context) *echo.HTTPError {
	message := "the requested resource could not be found"
	return echo.NewHTTPError(http.StatusNotFound, message)
}

func (app *Application) ErrMethodNotAllowed(ctx echo.Context) *echo.HTTPError {
	message := fmt.Sprintf("the %s method is not supported for this resource", ctx.Request().Method)
	return echo.NewHTTPError(http.StatusNotFound, message)
}

func (app *Application) ErrInternalServer(ctx echo.Context, err error) error {
	app.logError(ctx, err)
	message := "the server encountered a problem and could not process your request"
	return echo.NewHTTPError(http.StatusInternalServerError, message)
}

func (app *Application) ErrBadRequest(ctx echo.Context, err error) error {
	return echo.NewHTTPError(http.StatusBadRequest, err)
}

func (app *Application) ErrFailedValidation(ctx echo.Context, err error) error {
	message := "unable to process entity because some malformed value"
	return echo.NewHTTPError(http.StatusUnprocessableEntity, message).SetInternal(err)
}

func (app *Application) ErrInvalidCredentials(ctx echo.Context) error {
	msg := "invalid authentication credentials"
	return echo.NewHTTPError(http.StatusBadRequest, msg)
}

func (app *Application) ErrInvalidAuthenticationToken(ctx echo.Context) error {
	ctx.Request().Header.Set("WWW-Authenticate", "Bearer")

	msg := "invalid or missing authentication token"
	return echo.NewHTTPError(http.StatusBadRequest, msg)
}

// func ErrAuthenticationRequired(ctx *gin.Context) {
// 	msg := "you must be authenticated to access this resource"
// 	sendErr(ctx, http.StatusUnauthorized, msg, nil)
// }
