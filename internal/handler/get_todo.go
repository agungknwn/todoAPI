package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"github.com/waksun0x00/todoAPI/api"
	"github.com/waksun0x00/todoAPI/internal/tools"
)

func GetTodoDetails(w http.ResponseWriter, r *http.Request) {
	// var id = api.TodoParams{}
	// var decoder *schema.Decoder = schema.NewDecoder()
	var err error
	var database *tools.Todo

	// err = decoder.Decode(&id, r.URL.Query())

	// if err != nil {
	// 	log.Error(err)
	// 	api.InternalErrorHandler(w)
	// 	return
	// }

	id := chi.URLParam(r, "id")

	database, err = tools.GetTodoList(id)
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	var response = api.TodoResponse{
		Details: (*database),
		Code:    http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

}

func GetTodolist(w http.ResponseWriter, r *http.Request) {
	var err error
	var database *[]tools.Todo

	database, err = tools.GetTodo()
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}
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
