package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func writeJSON(w http.ResponseWriter, status int, data map[string]any) error {
	var buff bytes.Buffer
	err := json.NewEncoder(&buff).Encode(data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(buff.Bytes())

	return nil
}

func readJSON(w http.ResponseWriter, r *http.Request, dst any) error {
	err := json.NewDecoder(r.Body).Decode(dst)
	if err != nil {
		return err
	}
	
	return nil
}
