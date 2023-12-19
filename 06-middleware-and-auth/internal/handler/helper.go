package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func paramID(r *http.Request) (int64, error) {
	id, err := strconv.ParseInt(httprouter.ParamsFromContext(r.Context()).ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return -1, fmt.Errorf("invalid id parameter")
	}

	return int64(id), nil
}
