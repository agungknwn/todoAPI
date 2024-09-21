package handler

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/waksun0x00/todoAPI/api"
	"github.com/waksun0x00/todoAPI/internal/tools"
)

func (svc *APIservice) CreateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "Application/json")

	resp := &api.TodoResponse{}
	// Decode the JSON request body into the NewTodo struct
	defer json.NewEncoder(w).Encode(resp)

	var todo tools.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Invalid body", err) // handler for bad requests
		resp.Code = err.Error()
		return
	}

	todo.ID = uuid.New().String()

	repo := tools.TodoRepo{DBcollection: svc.MongoCollections}

	// insert Todo
	insertID, err := repo.InsertTodo(&todo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Insert todo failed", err)
		resp.Code = err.Error()
		return
	}

	resp.Data = todo.ID

	w.WriteHeader(http.StatusOK)
	log.Println("Todo inserted with id", insertID, todo)
}
