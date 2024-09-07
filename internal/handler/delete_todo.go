package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"github.com/waksun0x00/todoAPI/api"
	"github.com/waksun0x00/todoAPI/internal/tools"
)

func DeleteTodoDetails(w http.ResponseWriter, r *http.Request) {
	var err error

	id := chi.URLParam(r, "id")

	var database *[]tools.Todo
	database, err = tools.DeleteTodoDetails(id)
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	var response = api.TodoListResponse{
		TodoList: (*database),
		Code:     http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

}
