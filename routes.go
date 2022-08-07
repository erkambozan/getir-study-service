
package main

func (app *App) AddRoutes() {
    app.AddRoute("POST", "/in-memory", app.AddInMemoryItem)
    app.AddRoute("GET", "/in-memory/{key}", app.FindInMemoryItem)
    app.AddRoute("POST", "/records", app.FindRecords)
}