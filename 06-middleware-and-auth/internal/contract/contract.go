package contract

import "net/http"

type Contract struct {
	Error   ErrorResponse
	Success SuccessResponse
}

func New() Contract {
	return Contract{
		Error:   ErrorResponse{},
		Success: SuccessResponse{},
	}
}

// Write JSON
func (c Contract) Marshall(w http.ResponseWriter, status int, data map[string]any, headers http.Header) error {
	return writeJSON(w, status, data, headers)
}

// Read JSON
func (c Contract) Unmarshall(w http.ResponseWriter, r *http.Request, dst any) error {
	return readJSON(w, r, dst)
}
