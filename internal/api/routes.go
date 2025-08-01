package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (a *App) loadRoutes() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/", a.loadOrderRoutes)

	a.router = router
}



func (a *App) loadOrderRoutes(router chi.Router) {
	router.Get("/{key}", a.List)
	router.Get("/{key}/{id}", a.GetByID)
}
