package handler

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/waksun0x00/todoAPI/api"
	"github.com/waksun0x00/todoAPI/internal/tools"
)

func CreateTodolist(w http.ResponseWriter, r *http.Request) {
	var err error
	var database *[]tools.Todo

	// Decode the JSON request body into the NewTodo struct
	var NewTodo tools.Todo
	err = json.NewDecoder(r.Body).Decode(&NewTodo)
	if err != nil {
		log.Error(err)
		api.RequestErrorHandler(w, err) // handler for bad requests
		return
	}

	database, err = tools.CreateTodo(&NewTodo)
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
