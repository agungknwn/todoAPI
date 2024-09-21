package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"github.com/waksun0x00/todoAPI/api"
	"github.com/waksun0x00/todoAPI/internal/tools"
)

func (svc *APIservice) UpdateTodoByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "Application/json")

	resp := &api.TodoResponse{}
	// Decode the JSON request body into the NewTodo struct
	defer json.NewEncoder(w).Encode(resp)

	todoID := chi.URLParam(r, "id")
	if todoID == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("invalid todoID")
		resp.Code = "Invalid todo ID"
		return
	}

	var todo tools.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Invalid body", err) // handler for bad requests
		resp.Code = err.Error()
		return
	}

	todo.ID = todoID

	repo := tools.TodoRepo{DBcollection: svc.MongoCollections}

	// insert Todo
	count, err := repo.UpdateTodo(todoID, &todo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Update todo failed", err)
		resp.Code = err.Error()
		return
	}

	resp.Data = count

	w.WriteHeader(http.StatusOK)
}
