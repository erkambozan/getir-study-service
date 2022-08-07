package main

import (
	"encoding/json"
	"net/http"

	"getir-study-service/internal/dto"
	"getir-study-service/configs"


	"github.com/gorilla/mux"
)

type App struct {
	Router       *mux.Router
}

// NewApp creates a new instance of App
func NewApp() *App {
	return &App{
		Router: mux.NewRouter(),
	}
}

func (app *App) Initialize(){
	configs.ConnectDB()

	app.AddRoutes()
}

func (app *App) AddRoute(method string, route string, apiHandler func(w http.ResponseWriter, r *http.Request)) {
	app.Router.HandleFunc(route, apiHandler).Methods(method)
}

// RenderJSON creates a json response of the data
func (app *App) RenderJSON(writer http.ResponseWriter, status int, data interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	if data != nil {
		jsonData, err := json.Marshal(data)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		writer.Write(jsonData)
	}
}

// RenderErrorResponse render error response
func (app *App) RenderErrorResponse(writer http.ResponseWriter, httpStatus int, err error, message string) {
	writer.WriteHeader(httpStatus)
	app.RenderJSON(writer, httpStatus, dto.ErrorResponse{Status: httpStatus, Error: err, Message: message})
}