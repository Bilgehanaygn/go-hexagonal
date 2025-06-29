package router

import (
	"github.com/bilgehanaygn/urun/internal/category/infra/http/controller"
	"github.com/go-chi/chi/v5"
)

func Register(r *chi.Mux, c *controller.CategoryController) {
	r.Route("/category", func(r chi.Router) {
		r.Post("/", c.HandleCreate)
		r.Put("/", c.HandleUpdate)
		r.Get("/", c.HandleList)
		r.Get("/{id}", c.HandleGetById)
	})
}