package Routers

import (
	c "mongoDB-Golang/controllers"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/users", c.CreateUser).Methods("POST")
	router.HandleFunc("/users", c.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", c.GetUser).Methods("GET")
	router.HandleFunc("/users/{id}", c.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", c.DeleteUser).Methods("DELETE")

	return router

}
