package dto

import (
	"fmt"
	"net/http"

	"getir-study-service/models"
)

// RecordRequest request struct
type RecordRequest struct {
	StartDate   string `json:"startDate"`
	EndDate string `json:"endDate"`
	MinCount   int `json:"minCount"`
	MaxCount int `json:"maxCount"`
}

// RecordResponse response struct
type RecordResponse struct {
	Code   int `json:"code"`
	Status string `json:"status"`
	Records []*models.Record `json:"records"`
}

// ValidateRecord records request
func (request *RecordRequest) ValidateRecord() (int, error) {

	if request.StartDate == "" {
		return http.StatusBadRequest, fmt.Errorf("StartDate cannot be empty")
	}
	if request.EndDate == "" {
		return http.StatusBadRequest, fmt.Errorf("EndDate cannot be empty")
	}
	if request.MinCount == 0 {
		return http.StatusBadRequest, fmt.Errorf("Value cannot be empty")
	}
	if request.MaxCount == 0 {
		return http.StatusBadRequest, fmt.Errorf("Value cannot be empty")
	}

	return http.StatusOK, nil
}