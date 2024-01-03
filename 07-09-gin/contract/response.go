package contract

type ApiError struct {
	StatusCode string `json:"status_code"`
	Details    any    `json:"details,omitempty"`
}

type APIResponse struct {
	Success bool      `json:"success"`
	Message string    `json:"message"`
	Data    any       `json:"data"`
	Error   *ApiError `json:"error"`
}
