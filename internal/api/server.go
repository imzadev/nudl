package api

import (
	"fmt"
	"net/http"
)

type App struct {
	router  http.Handler
	content map[string]any
	port    string
}

func New(port string, data map[string]any) *App {
	app := &App{
		content: data,
		port:    port,
	}

	app.loadRoutes()

	return app
}

func (a *App) Start() {
	server := &http.Server{
		Addr:    ":" + a.port,
		Handler: a.router,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("failed to listen to server", err)
	}

}
