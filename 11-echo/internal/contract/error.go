package contract

type ErrorResponse struct {
	StatusText string `json:"status_text"`
	Message    any    `json:"message"`
	Details    any    `json:"details,omitempty"`
}
