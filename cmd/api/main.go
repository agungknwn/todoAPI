package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/waksun0x00/todoAPI/internal/handler"
)

func main() {
	var router *chi.Mux = chi.NewRouter()
	handler.Handler(router)

	fmt.Println("Starting GO RESTful API...")

	err := http.ListenAndServe("localhost:6666", router)

	if err != nil {
		log.Fatal(err)
	}
}
