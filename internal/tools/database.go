package tools

import "go.mongodb.org/mongo-driver/mongo"

type Todo struct {
	ID     string `json:"id" bson:"id"`
	Desc   string `json:"description" bson:"description"`
	Status string `json:"status" bson:"status"`
}

type TodoRepo struct {
	DBcollection *mongo.Collection
}
