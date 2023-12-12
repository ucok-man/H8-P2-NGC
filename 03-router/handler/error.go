package handler

import (
	"fmt"
	"net/http"
)

func errorResponse(w http.ResponseWriter, status int, msg string) {
	err := writeJSON(w, status, map[string]any{"error": msg})
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
	}
}

func errorServer(w http.ResponseWriter, err error) {
	fmt.Println("err:", err)
	errorResponse(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
}

func errorClient(w http.ResponseWriter, status int, msg string) {
	errorResponse(w, status, msg)
}
