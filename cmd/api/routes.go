package main
import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/status",app.statusHandler)
	router.Get("/v1/movie/{movid}",app.getMovie)
	router.Get("/v1/movies",app.getAllMovies)
	return router
}