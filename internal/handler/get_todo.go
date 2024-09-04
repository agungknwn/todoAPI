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
	// var params = api.TodoParams{}
	// var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	// err = decoder.Decode(&params, r.URL.Query())

	// if err != nil {
	// 	log.Error(err)
	// 	api.InternalErrorHandler(w)
	// 	return
	// }

	id := chi.URLParam(r, "id")

	var database tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	var todoDetails *tools.Todo = database.GetTodoList(id)
	if todoDetails == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var response = api.TodoResponse{
		Details: (*todoDetails),
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
	var database tools.DatabaseInterface

	database, err = tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	var todoList *[]tools.Todo = database.GetTodo()
	if *todoList == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
	}

	var response = api.TodoListResponse{
		TodoList: (*todoList),
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
