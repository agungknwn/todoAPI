package tools

import (
	"fmt"
	"time"
)

var mockTodoList = []Todo{
	{ID: "123ABC", Desc: "Cuci Baju", Status: "On Going"},
	{ID: "234BCA", Desc: "Mandi", Status: "On Going"},
	{ID: "345DEF", Desc: "Makan", Status: "On Going"},
}

func GetTodoList(id string) (*Todo, error) {
	// DB call simulations
	time.Sleep(time.Second * 1)
	var err error

	var clientData = Todo{}
	for i := range mockTodoList {
		if id == mockTodoList[i].ID {
			clientData = mockTodoList[i]
			err = nil
			break
		}
	}

	return &clientData, err
}

func GetTodo() (*[]Todo, error) {
	var err error
	time.Sleep(time.Second * 1)
	err = nil
	return &mockTodoList, err
}

func CreateTodo(todo *Todo) (*[]Todo, error) {
	time.Sleep(time.Second * 1)
	var err error = nil

	mockTodoList = append(mockTodoList, *todo)

	fmt.Println("Successfully Add new Task")

	return &mockTodoList, err
}

func UpdateTodoDetails(id string, newTodo string, newStatus string) (*[]Todo, error) {
	time.Sleep(time.Second * 1)
	// var clientData []Todo = mockTodoList
	var err error = nil

	for i := range mockTodoList {
		if id == mockTodoList[i].ID {
			mockTodoList[i].Desc = newTodo
			mockTodoList[i].Status = newStatus
			break
		}
	}

	fmt.Println("Successfully Update Task")

	return &mockTodoList, err
}

func DeleteTodoDetails(id string) (*[]Todo, error) {
	time.Sleep(time.Second * 1)

	// var clientData []Todo = mockTodoList
	var err error = nil

	for i := range mockTodoList {
		if id == mockTodoList[i].ID {
			mockTodoList = append(mockTodoList[:i], mockTodoList[i+1:]...)
			break
		}
	}

	fmt.Println("Successfully Delete Task")

	return &mockTodoList, err
}
