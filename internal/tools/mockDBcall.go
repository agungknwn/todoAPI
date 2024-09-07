package tools

import (
	"fmt"
	"time"
)

type mockDB struct{}

var mockTodoList = []Todo{
	{ID: "123ABC", Desc: "Cuci Baju", Status: "On Going"},
	{ID: "234BCA", Desc: "Mandi", Status: "On Going"},
	{ID: "345DEF", Desc: "Makan", Status: "On Going"},
}

func (d *mockDB) GetTodoList(id string) *Todo {
	// DB call simulations
	time.Sleep(time.Second * 1)

	var clientData = Todo{}
	var database = []Todo{}
	database = mockTodoList
	for i := range database {
		if id == database[i].ID {
			clientData = database[i]
			break
		}
	}

	return &clientData
}

func (d *mockDB) GetTodo() *[]Todo {
	time.Sleep(time.Second * 1)
	return &mockTodoList
}

func (d *mockDB) CreateTodo(todo *Todo) *[]Todo {
	time.Sleep(time.Second * 1)

	mockTodoList = append(mockTodoList, *todo)

	fmt.Println("Successfully Add new Task")

	return &mockTodoList
}

func (d *mockDB) UpdateTodoDetails(id string, newTodo string, newStatus string) *[]Todo {
	time.Sleep(time.Second * 1)
	// var clientData []Todo = mockTodoList

	for i := range mockTodoList {
		if id == mockTodoList[i].ID {
			mockTodoList[i].Desc = newTodo
			mockTodoList[i].Status = newStatus
			break
		}
	}

	fmt.Println("Successfully Update Task")

	return &mockTodoList
}

func (d *mockDB) DeleteTodoDetails(id string) *[]Todo {
	time.Sleep(time.Second * 1)

	// var clientData []Todo = mockTodoList

	for i := range mockTodoList {
		if id == mockTodoList[i].ID {
			mockTodoList = append(mockTodoList[:i], mockTodoList[i+1:]...)
			break
		}
	}

	return &mockTodoList
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
