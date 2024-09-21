package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"github.com/waksun0x00/todoAPI/internal/handler"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoClient *mongo.Client

func init() {
	// load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("env Load error", err)
	}

	log.Print(".env file Loaded")

	mongoClient, err = mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGO_URI")))

	if err != nil {
		log.Fatal("Connection error", err)
	}

	err = mongoClient.Ping(context.Background(), readpref.Primary())

	if err != nil {
		log.Fatal("Ping Failed ", err)
	}

	log.Print("DB connected")
}

func main() {
	defer mongoClient.Disconnect(context.Background())

	var router *chi.Mux = chi.NewRouter()
	handler.Handler(mongoClient, router)

	fmt.Println("Starting GO RESTful API...")

	err := http.ListenAndServe("localhost:8000", router)

	if err != nil {
		log.Fatal(err)
	}
}
