package main

import (
	"log"
	"mongoDB-Golang/Routers"
	"mongoDB-Golang/controllers"
	"mongoDB-Golang/modules"
	"net/http"
)

func main() {
	modules.ConnectDatabase()
	controllers.InitCollection()
	router := Routers.SetupRouter()

	userCollection := modules.GetCollection("mongoDB-GO", "Users")
	modules.CreateInitialAdmin(userCollection)

	log.Fatal(http.ListenAndServe(":8000", router))
}
