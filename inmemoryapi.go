package main

import (
	"encoding/json"
	"io"
	"net/http"
	"github.com/gorilla/mux"
	
	"getir-study-service/internal/dto"
	"getir-study-service/internal/store"
)

// AddInMemoryItem adds news item to in-memory and creates a json response of the data
func (app *App) AddInMemoryItem(writer http.ResponseWriter, req *http.Request) {

	var reqBody []byte
	var err error

	// read the request body
	if reqBody, err = io.ReadAll(req.Body); err != nil {
		app.RenderErrorResponse(writer, http.StatusBadRequest, err, "Failed to read request body")
		return
	}

	// serialize to json
	var request dto.InMemoryRequest
	if err = json.Unmarshal(reqBody, &request); err != nil {
		app.RenderErrorResponse(writer, http.StatusInternalServerError, err, "Failed to convert json code")
		return
	}

	// validate fields
	if status, errValidate := request.ValidateInmemory(); errValidate != nil {
		app.RenderErrorResponse(writer, status, errValidate, "Validation error")
		return
	}

	// insert message
	response, errResponse := store.CreateInMemoryItem(request)
	if errResponse != nil {
		app.RenderErrorResponse(writer, errResponse.Status, errResponse.Error, errResponse.Message)
		return
	}

	// render output
	app.RenderJSON(writer, http.StatusOK, response)

}

// GetInMemoryItem get in memory item from in-memory and creates a json response of the data
func (app *App) FindInMemoryItem(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	key := params["key"]

	response, errResponse := store.FindInMemoryItem(key)
	if errResponse != nil {
		app.RenderErrorResponse(writer, errResponse.Status, errResponse.Error, errResponse.Message)
		return
	}

	// render output
	app.RenderJSON(writer, http.StatusOK, response)
}