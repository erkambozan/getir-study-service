package main

import (
    "log"
    "net/http"
    "os"
    "fmt"
)

func main() {
    app := NewApp()

    app.Initialize()

    port := os.Getenv("PORT")

    if port == "" {
        port = "3000"
    }
    fmt.Printf("Port : %s", port)
    log.Fatal(http.ListenAndServe(":"+port, app.Router))

}