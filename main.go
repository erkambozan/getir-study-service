package main

import (
    "log"
    "net/http"

)

func main() {
    app := NewApp()

    app.Initialize()

    port, err := os.Getenv("PORT")

    log.Fatal(http.ListenAndServe(":"+port, app.Router))

}