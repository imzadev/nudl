package api

import (
	"net/http"
	"fmt"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/imzadev/nudl/internal/config"
	"github.com/imzadev/nudl/internal/jsonloader"
)

func (a *App) loadRoutes() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/api", a.loadOrderRoutes)

	a.router = router
}



func (a *App) loadOrderRoutes(router chi.Router) {
	jsonPath := "/home/imza/coding/nudl/datatest/test.json" 
	Data,err := jsonloader.LoadAndParse(jsonPath)
	if err != nil {
	 fmt.Println(err)
	}
	content := Content{
		data: config.NewData(Data),
	}
	
	router.Get("/{key}", content.List)
	router.Get("/{key}/{id}", content.GetByID)
}
