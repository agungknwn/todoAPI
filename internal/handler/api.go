package handler

import (
	"os"

	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/mongo"
	// chimiddle "github.com/go-chi/chi/middleware"
)

type APIservice struct {
	MongoCollections *mongo.Collection
}

func Handler(mongoClient *mongo.Client, r *chi.Mux) {

	coll := mongoClient.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("COLLECTION_NAME"))

	todoAPI := APIservice{MongoCollections: coll}
	// global middleware
	// r.Use(chimiddle.StripSlashes)
	r.Route("/TodoList", func(router chi.Router) {
		router.Get("/", todoAPI.GetTodolist)
		router.Get("/{id}", todoAPI.GetTodoID)
		router.Post("/", todoAPI.CreateTodo)
		router.Put("/{id}", todoAPI.UpdateTodoByID)
		router.Delete("/{id}", todoAPI.DeleteTodoByID)
	})
}
