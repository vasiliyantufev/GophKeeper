package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/vasiliyantufev/gophkeeper/internal/server/api/rest"
)

// Route - setting service routes
func Route(s *resthandler.Handler) *chi.Mux {
	r := chi.NewRouter()

	r.Route("/api", func(r chi.Router) {
		r.Get("/", s.IndexHandler)
		r.Route("/user", func(r chi.Router) {
			r.Post("/block", s.UserBlock)
			r.Post("/unblock", s.UserUnblock)
		})
		r.Route("/token", func(r chi.Router) {
			r.Get("/{userId}", s.TokenGetList)
			r.Post("/block", s.TokenBlock)
		})
	})
	return r
}
