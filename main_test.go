package main

import(
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var app *App


func TestMain(m *testing.M) {
	app = NewApp()
	fmt.Println("Running application...")
	app.Initialize()
	m.Run()
}

func CheckResponseCode(t *testing.T, expected int, response *httptest.ResponseRecorder) {
	if expected != response.Code {
		t.Fatalf("Expected response code %d. Got %d and Body:%s\n", expected, response.Code, response.Body)
	}
}

func ExecuteRequest(request *http.Request) *httptest.ResponseRecorder {
	request.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	app.Router.ServeHTTP(recorder, request)
	return recorder
}