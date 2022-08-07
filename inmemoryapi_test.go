package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
	"fmt"

	"getir-study-service/internal/dto"
)

func TestAddInMemoryItem(t *testing.T) {
	inMemoryRequest := dto.InMemoryRequest{
		Key: "active-tabs",
		Value: "getir",
	}

	jsonString, _ := json.Marshal(inMemoryRequest)
	request, _ := http.NewRequest("POST", "/in-memory", bytes.NewReader(jsonString))
	response := ExecuteRequest(request)
	CheckResponseCode(t, http.StatusOK, response)
}

func TestFindInMemoryItem(t *testing.T) {
	expectedResponse := dto.InMemoryResponse{
		Key: "active-tabs",
		Value: "getir",
	}

	key :=  "active-tabs"
	request, _ := http.NewRequest("GET", fmt.Sprintf("/in-memory/%s", key), nil)

	response := ExecuteRequest(request)

	actualResponse := dto.InMemoryResponse{}
	CheckResponseCode(t, http.StatusOK, response)
	json.Unmarshal(response.Body.Bytes(), &actualResponse)

	if actualResponse.Key != expectedResponse.Key {
		t.Fatalf("Expected key to be '%s'. Got '%s'", actualResponse.Key, expectedResponse.Key)
	}

	if actualResponse.Value != expectedResponse.Value {
		t.Fatalf("Expected value to be '%s'. Got '%s'", actualResponse.Value, expectedResponse.Value)
	}
}