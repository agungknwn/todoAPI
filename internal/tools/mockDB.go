package tools

import (
	"time"
)

type mockDB struct{}

var mockTodoList = []Todo{
	{ID: "1", Desc: "Cuci Baju", Status: "On Going"},
	{ID: "2", Desc: "Mandi", Status: "On Going"},
	{ID: "3", Desc: "Makan", Status: "On Going"},
}

func (d *mockDB) GetTodoList(id int) *Todo {
	// DB call simulations
	time.Sleep(time.Second * 1)

	var clientData = Todo{}
	clientData = mockTodoList[id]

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}

func GetTodo() []Todo {
	time.Sleep(time.Second * 1)

	return mockTodoList
}
