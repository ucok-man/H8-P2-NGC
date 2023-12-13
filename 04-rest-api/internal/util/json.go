package util

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func (u Utility) WriteJSON(w http.ResponseWriter, status int, data map[string]any, headers http.Header) error {
	var buff bytes.Buffer
	err := json.NewEncoder(&buff).Encode(data)
	if err != nil {
		return err
	}

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(buff.Bytes())

	return nil
}

func (u Utility) ReadJSON(w http.ResponseWriter, r *http.Request, dst any) error {
	err := json.NewDecoder(r.Body).Decode(dst)
	if err != nil {
		return err
	}

	return nil
}
