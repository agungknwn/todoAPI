package tools

import (
	log "github.com/sirupsen/logrus"
)

type Todo struct {
	ID     string `json:"id"`
	Desc   string `json:"description"`
	Status string `json:"status"`
}

type DatabaseInterface interface {
	GetTodoList(id string) *Todo
	GetTodo() *[]Todo
	CreateTodo(todo *Todo) *[]Todo
	UpdateTodoDetails(id string, newTodo string, newStatus string) *[]Todo
	DeleteTodoDetails(id string) *[]Todo
	SetupDatabase() error
}

func NewDatabase() (DatabaseInterface, error) {
	var database DatabaseInterface = &mockDB{}

	var err error = database.SetupDatabase()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return database, nil
}
