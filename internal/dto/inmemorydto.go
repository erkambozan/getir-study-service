package dto

import (
	"fmt"
	"net/http"
)

// InMemoryRequest request struct
type InMemoryRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// InMemoryResponse response struct
type InMemoryResponse struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// ValidateInMemory validates request
func (request *InMemoryRequest) ValidateInmemory() (int, error) {

	if request.Key == "" {
		return http.StatusBadRequest, fmt.Errorf("Key cannot be empty")
	}
	if request.Value == "" {
		return http.StatusBadRequest, fmt.Errorf("Value cannot be empty")
	}

	return http.StatusOK, nil
}