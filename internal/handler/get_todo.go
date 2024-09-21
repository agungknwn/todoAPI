package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"github.com/waksun0x00/todoAPI/api"
	"github.com/waksun0x00/todoAPI/internal/tools"
)

func (svc *APIservice) GetTodoID(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "Application/json")

	resp := &api.TodoResponse{}

	defer json.NewEncoder(w).Encode(resp)

	todoID := chi.URLParam(r, "id")

	repo := tools.TodoRepo{DBcollection: svc.MongoCollections}

	todo, err := repo.FindTodoID(todoID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Get todo failed", err)
		resp.Code = err.Error()
		return
	}
	resp.Data = todo
	w.WriteHeader(http.StatusOK)
}

func (svc *APIservice) GetTodolist(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "Application/json")

	resp := &api.TodoResponse{}

	defer json.NewEncoder(w).Encode(resp)

	repo := tools.TodoRepo{DBcollection: svc.MongoCollections}

	todoList, err := repo.FindTodoList()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Get todo failed", err)
		resp.Code = err.Error()
		return
	}

	resp.Data = todoList
	w.WriteHeader(http.StatusOK)

}
