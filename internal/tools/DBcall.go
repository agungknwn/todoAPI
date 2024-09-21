package tools

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *TodoRepo) FindTodoID(todoID string) (*Todo, error) {
	var todo Todo
	// DB call
	err := r.DBcollection.FindOne(context.Background(), bson.D{{Key: "id", Value: todoID}}).Decode(&todo)

	if err != nil {
		return nil, err
	}

	return &todo, err
}

func (r *TodoRepo) FindTodoList() ([]Todo, error) {
	resp, err := r.DBcollection.Find(context.Background(), bson.D{})

	if err != nil {
		return nil, err
	}

	var todoList []Todo
	err = resp.All(context.Background(), &todoList)

	if err != nil {
		return nil, err
	}

	return todoList, err
}

func (r *TodoRepo) InsertTodo(todo *Todo) (interface{}, error) {

	resp, err := r.DBcollection.InsertOne(context.Background(), todo)

	if err != nil {
		return nil, err
	}

	return resp, err
}

func (r *TodoRepo) UpdateTodo(todoID string, newTodo *Todo) (int64, error) {
	resp, err := r.DBcollection.UpdateOne(context.Background(),
		bson.D{{Key: "id", Value: todoID}},
		bson.D{{Key: "$set", Value: *newTodo}})

	if err != nil {
		return 0, err
	}

	return resp.ModifiedCount, err
}

func (r *TodoRepo) DeleteTodo(todoID string) (int64, error) {
	resp, err := r.DBcollection.DeleteOne(context.Background(), bson.D{{Key: "id", Value: todoID}})

	if err != nil {
		return 0, err
	}

	fmt.Println("Successfully Delete Task")

	return resp.DeletedCount, err
}
