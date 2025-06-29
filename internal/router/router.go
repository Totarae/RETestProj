package router

import (
	"awesomeProject10/internal/config"
	"awesomeProject10/internal/handlers"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func NewRouter(conf *config.Config) http.Handler {
	r := chi.NewRouter()
	//r.Use(middleware.Logger)

	handler := handlers.NewHandler(conf)

	r.Get("/", handler.ServeIndex)         // UI
	r.Post("/calculate", handler.Optimize) // API

	// Static files ( CSS/JS)
	r.Handle("/web/*", http.StripPrefix("/web/", http.FileServer(http.Dir("web/web"))))

	return r
}
