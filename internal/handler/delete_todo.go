package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"github.com/waksun0x00/todoAPI/api"
	"github.com/waksun0x00/todoAPI/internal/tools"
)

func (svc *APIservice) DeleteTodoByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "Application/json")

	resp := &api.TodoResponse{}
	// Decode the JSON request body into the NewTodo struct
	defer json.NewEncoder(w).Encode(resp)

	todoID := chi.URLParam(r, "id")

	repo := tools.TodoRepo{DBcollection: svc.MongoCollections}

	// insert Todo
	count, err := repo.DeleteTodo(todoID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error :", err)
		resp.Code = err.Error()
		return
	}

	resp.Data = count

	w.WriteHeader(http.StatusOK)
}
