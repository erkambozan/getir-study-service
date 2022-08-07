package dto

// ErrorResponse a wrapper for error response
type ErrorResponse struct {
	Status  int    `json:"status"`
	Error   error  `json:"error"`
	Message string `json:"message"`
}