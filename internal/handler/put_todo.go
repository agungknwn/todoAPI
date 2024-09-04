package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"github.com/waksun0x00/todoAPI/api"
	"github.com/waksun0x00/todoAPI/internal/tools"
)

func UpdateTodoDetails(w http.ResponseWriter, r *http.Request) {
	var err error
	var database tools.DatabaseInterface

	database, err = tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	// Decode the JSON request body into the NewTodo struct
	id := chi.URLParam(r, "id")

	var updateTodo struct {
		NewTodo   string `json:"description"`
		NewStatus string `json:"status"`
	}

	err = json.NewDecoder(r.Body).Decode(&updateTodo)
	if err != nil {
		log.Error(err)
		api.RequestErrorHandler(w, err)
		return
	}

	var UpdatedTodoList *[]tools.Todo = database.UpdateTodoDetails(id, updateTodo.NewTodo, updateTodo.NewStatus)
	if UpdatedTodoList == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
	}

	var response = api.TodoListResponse{
		TodoList: (*UpdatedTodoList),
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
