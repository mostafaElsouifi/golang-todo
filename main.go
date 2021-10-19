package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// routes
	router.HandleFunc("/", getAllTodos).Methods("GET")
	router.HandleFunc("/todo/{id}", getTodo).Methods("GET")
	router.HandleFunc("/todo", createNewTodo).Methods("POST")
	//router.HandleFunc("/todo/{id}", updateTodo).Methods("PUT")
	router.HandleFunc("/todo/{id}", deleteTodo).Methods("DELETE")
	http.ListenAndServe(":3000", router)
}
