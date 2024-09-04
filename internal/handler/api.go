package handler

import (
	"github.com/go-chi/chi"
	// chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) {
	// global middleware
	// r.Use(chimiddle.StripSlashes)
	r.Route("/TodoList", func(router chi.Router) {

		router.Get("/", GetTodolist)
		router.Get("/{id}", GetTodoDetails)
		router.Post("/", CreateTodolist)
		router.Put("/{id}", UpdateTodoDetails)
		router.Delete("/{id}", DeleteTodoDetails)
	})
}
