package contract

import (
	"net/http"

	orderedmap "github.com/wk8/go-ordered-map/v2"
)

type SuccessResponse struct{}

func (res SuccessResponse) sendSuccess(w http.ResponseWriter, status int, msg, key string, data any) {
	response := orderedmap.New[string, any]()
	response.Set("code", http.StatusText(status))
	response.Set("message", msg)
	response.Set(key, data)

	err := writeJSON(w, status, map[string]any{"data": response}, nil)
	if err != nil {
		w.WriteHeader(500)
	}
}

func (res SuccessResponse) OK(w http.ResponseWriter, msg, key string, data any) {
	res.sendSuccess(w, http.StatusOK, msg, key, data)
}

func (res SuccessResponse) Created(w http.ResponseWriter, msg, key string, data any) {
	res.sendSuccess(w, http.StatusCreated, msg, key, data)
}
