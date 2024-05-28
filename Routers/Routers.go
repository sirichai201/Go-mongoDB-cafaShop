package Routers

import (
	"mongoDB-Golang/controllers"
	"mongoDB-Golang/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	//login_register
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.Handle("/users", (http.HandlerFunc(controllers.CreateUser))).Methods("POST")

	//testrole
	router.Handle("/admin", middlewares.AuthenticationMiddleware(middlewares.AdminMiddleware(http.HandlerFunc(controllers.AdminHandler)))).Methods("GET")
	router.Handle("/user", middlewares.AuthenticationMiddleware(http.HandlerFunc(controllers.UserHandler))).Methods("GET")

	//user
	router.Handle("/users", middlewares.AuthenticationMiddleware(http.HandlerFunc(controllers.GetUsers))).Methods("GET")
	router.Handle("/users/{id}", middlewares.AuthenticationMiddleware(http.HandlerFunc(controllers.GetUser))).Methods("GET")
	router.Handle("/users/{id}", middlewares.AuthenticationMiddleware(http.HandlerFunc(controllers.UpdateUser))).Methods("PUT")
	router.Handle("/users/{id}", middlewares.AuthenticationMiddleware(http.HandlerFunc(controllers.DeleteUser))).Methods("DELETE")

	//product
	router.Handle("/products", middlewares.AuthenticationMiddleware(http.HandlerFunc(controllers.GetProducts))).Methods("GET")
	router.Handle("/products", middlewares.AuthenticationMiddleware(http.HandlerFunc(controllers.CreateProduct))).Methods("POST")
	router.Handle("/products/{id}", middlewares.AuthenticationMiddleware(http.HandlerFunc(controllers.UpdateProduct))).Methods("PUT")
	router.Handle("/products/{id}", middlewares.AuthenticationMiddleware(http.HandlerFunc(controllers.DeleteProduct))).Methods("DELETE")

	return router
}
