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
	var database tools.DatabaseInterface

	database, err = tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	// Decode the JSON request body into the NewTodo struct
	var NewTodo tools.Todo
	err = json.NewDecoder(r.Body).Decode(&NewTodo)
	if err != nil {
		log.Error(err)
		api.RequestErrorHandler(w, err) // Assuming you have a handler for bad requests
		return
	}

	var NewTodoList *[]tools.Todo = database.CreateTodo(&NewTodo)
	if NewTodoList == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
	}

	var response = api.TodoListResponse{
		TodoList: (*NewTodoList),
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
