package main

import (
    "log"
    "net/http"

)

func main() {
    app := NewApp()

    app.Initialize()

    log.Fatal(http.ListenAndServe(":4004", app.Router))

}