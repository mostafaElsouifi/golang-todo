package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo/options"
)

// create collection
var todosCollection = db().Database("todo-app").Collection("todos")

// struct model
type Todo struct {
	ID   primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Item string             `json:"item,omitempty" bson:"item,omitempty"`
}

func createNewTodo(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	var todo Todo
	err := json.NewDecoder(req.Body).Decode(&todo)
	if err != nil {
		fmt.Println(err)
	}
	todoInserted, err := todosCollection.InsertOne(context.TODO(), todo)
	if err != nil {
		log.Fatal(err)

	}
	json.NewEncoder(res).Encode(todoInserted)
}

// get all todos
func getAllTodos(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	var todos []Todo
	cursor, err := todosCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var todo Todo
		cursor.Decode(&todo)
		todos = append(todos, todo)
	}
	json.NewEncoder(res).Encode(todos)
}

// get todo

func getTodo(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	params := mux.Vars(req)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var todo Todo
	err := todosCollection.FindOne(context.TODO(), Todo{ID: id}).Decode(&todo)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(res).Encode(todo)
}

// delete todo

func deleteTodo(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "application/json")
	params := mux.Vars(req)
	_id, _ := primitive.ObjectIDFromHex(params["id"])
	opts := options.Delete().SetCollation(&options.Collation{}) // need to understand
	deletedTodo, err := todosCollection.DeleteOne(context.TODO(), bson.D{{"_id", _id}}, opts)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(res).Encode(deletedTodo)
}

// update todo
