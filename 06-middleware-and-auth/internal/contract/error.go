package contract

import (
	"fmt"
	"net/http"

	orderedmap "github.com/wk8/go-ordered-map/v2"
)

type ErrorResponse struct{}

// for data if you want to NOT marshal them you must use nil. other than that will attemp
// to marshal the value.
func (res ErrorResponse) sendErr(w http.ResponseWriter, status int, msg string, data any) {
	response := orderedmap.New[string, any]()
	response.Set("code", http.StatusText(status))
	response.Set("message", msg)
	if data != nil {
		response.Set("details", data)
	}

	err := writeJSON(w, status, map[string]any{"error": response}, nil)
	if err != nil {
		w.WriteHeader(500)
	}
}

func (res ErrorResponse) InternalServer(w http.ResponseWriter, err error) {
	fmt.Println("err:", err)
	msg := "the server encountered a problem and could not process your request"
	res.sendErr(w, http.StatusInternalServerError, msg, nil)
}

func (res ErrorResponse) NotFound(w http.ResponseWriter, r *http.Request) {
	msg := "the requested resource could not be found"
	res.sendErr(w, http.StatusNotFound, msg, nil)
}

func (res ErrorResponse) MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	msg := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	res.sendErr(w, http.StatusMethodNotAllowed, msg, nil)
}

func (res ErrorResponse) BadRequest(w http.ResponseWriter, r *http.Request, err error) {
	res.sendErr(w, http.StatusBadRequest, err.Error(), nil)
}

func (res ErrorResponse) FailedValidation(w http.ResponseWriter, errors map[string]string) {
	msg := "unable to process entity because some malformed value"
	res.sendErr(w, http.StatusUnprocessableEntity, msg, errors)
}

func (res *ErrorResponse) InvalidCredentials(w http.ResponseWriter, r *http.Request) {
	message := "invalid authentication credentials"
	res.sendErr(w, http.StatusUnauthorized, message, nil)
}

func (res ErrorResponse) InvalidAuthenticationToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("WWW-Authenticate", "Bearer")

	message := "invalid or missing authentication token"
	res.sendErr(w, http.StatusUnauthorized, message, nil)
}

func (res *ErrorResponse) AuthenticationRequired(w http.ResponseWriter, r *http.Request) {
	message := "you must be authenticated to access this resource"
	res.sendErr(w, http.StatusUnauthorized, message, nil)
}

func (res *ErrorResponse) NotPermittedResource(w http.ResponseWriter, r *http.Request) {
	message := "your user account doesn't have the necessary permissions to access this resource"
	res.sendErr(w, http.StatusForbidden, message, nil)
}
