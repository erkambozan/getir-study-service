package store

import (
	"net/http"
    "os"
	"encoding/json"
	"io"

	"getir-study-service/internal/dto"
	"getir-study-service/models"
)

func CreateInMemoryItem(request dto.InMemoryRequest) (*dto.InMemoryResponse, *dto.ErrorResponse) {
	return &dto.InMemoryResponse{
		Key:  request.Key,
		Value: request.Value,
	}, nil
}

func FindInMemoryItem(key string) (*dto.InMemoryResponse, *dto.ErrorResponse) {
	inMemoryJson, err := os.Open("inmemory.json")

	if err != nil {
		return nil, &dto.ErrorResponse{Status: http.StatusInternalServerError, Error: err, Message: "Failed to read in memory file"}
	}

	var jsonBody []byte
	jsonBody, err = io.ReadAll(inMemoryJson)

	if err != nil {
		return nil, &dto.ErrorResponse{Status: http.StatusInternalServerError, Error: err, Message: "Failed to read request body"}
	}

	// serialize to json
	var model models.InMemory
	err = json.Unmarshal(jsonBody, &model)
	if err != nil {
		return nil, &dto.ErrorResponse{Status: http.StatusInternalServerError, Error: err, Message: "Failed to convert json code"}
	}

	if key != model.Key {
		return nil, &dto.ErrorResponse{Status: http.StatusInternalServerError, Error: err, Message: "Couldn't find any data"}
	}

	return &dto.InMemoryResponse{
		Key:  model.Key,
		Value: model.Value,
	}, nil
}
