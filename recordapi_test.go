package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"reflect"

	"getir-study-service/internal/dto"
	"getir-study-service/models"
)

func TestFindRecords(t *testing.T) {
	recordRequest := dto.RecordRequest{
			StartDate: "2016-01-26",
			EndDate: "2018-02-02",
			MinCount: 2700,
			MaxCount: 3000,
	}

	expectedRecords := []*models.Record{
		{
            Key: "TAKwGc6Jr4i8Z487",
            CreatedAt: "2017-01-28T01:22:14.398Z",
            TotalCount: 2800,
        },
        {
            Key: "NAeQ8eX7e5TEg7oH",
            CreatedAt: "2017-01-27T08:19:14.135Z",
            TotalCount: 2900,
        },
	}

	expectedResponse := dto.RecordResponse{
		Code: 200,
		Status: "OK",
		Records: expectedRecords,
	}

	var actualResponse dto.RecordResponse

	jsonString, _ := json.Marshal(recordRequest)
	request, _ := http.NewRequest("POST", "/records", bytes.NewReader(jsonString))
	response := ExecuteRequest(request)
	CheckResponseCode(t, http.StatusOK, response)

	json.Unmarshal(response.Body.Bytes(), &actualResponse)

	if !reflect.DeepEqual(actualResponse.Records, expectedResponse.Records) {
		t.Fatal("Expected Records are not same")
	}
}