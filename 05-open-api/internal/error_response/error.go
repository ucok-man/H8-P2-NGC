package error_response

import (
	"fmt"
	"net/http"
)

type JSONReadWrite interface {
	WriteJSON(w http.ResponseWriter, status int, data map[string]any, headers http.Header) error
	ReadJSON(w http.ResponseWriter, r *http.Request, dst any) error
}

type Error struct {
	json JSONReadWrite
}

func New(json JSONReadWrite) *Error {
	return &Error{
		json: json,
	}
}

func (e Error) errorResponse(w http.ResponseWriter, status int, msg string, details map[string]string) {
	data := ErrorContract{
		Code:    http.StatusText(status),
		Message: msg,
		Details: details,
	}

	err := e.json.WriteJSON(w, status, map[string]any{"error": data}, nil)
	if err != nil {
		w.WriteHeader(500)
	}
}

func (e Error) InternalServer(w http.ResponseWriter, err error) {
	fmt.Println("err:", err)
	msg := "the server encountered a problem and could not process your request"
	e.errorResponse(w, http.StatusInternalServerError, msg, nil)
}

func (e *Error) NotFound(w http.ResponseWriter, r *http.Request) {
	msg := "the requested resource could not be found"
	e.errorResponse(w, http.StatusNotFound, msg, nil)
}

func (e Error) MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	msg := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	e.errorResponse(w, http.StatusMethodNotAllowed, msg, nil)
}

func (e Error) BadRequest(w http.ResponseWriter, r *http.Request, err error) {
	e.errorResponse(w, http.StatusBadRequest, err.Error(), nil)
}

func (e Error) FailedValidation(w http.ResponseWriter, errors map[string]string) {
	msg := "unable to process entity because some malformed value"
	e.errorResponse(w, http.StatusUnprocessableEntity, msg, errors)
}
