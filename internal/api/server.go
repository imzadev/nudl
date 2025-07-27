package api

import (
	"fmt"
	"net/http"
)

type App struct {
	router http.Handler
}

func New() *App {
	app := &App{}

	app.loadRoutes()

	return app
}

func (a *App) Start(){
	server := &http.Server{
		Addr: ":3000",
		Handler: a.router,
	}
	
	fmt.Println("Starting server")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("failed to listen to server", err)
	}


}
