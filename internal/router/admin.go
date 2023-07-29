package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewAdminRouter() http.Handler {
	r := chi.NewRouter()

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/foos", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(200)
			w.Write([]byte("admin api foo"))
		})
	})
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("admin root"))
	})

	return r
}
