package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jahzielcabrera/golang-restapi/db"
	"github.com/jahzielcabrera/golang-restapi/models"
	"github.com/jahzielcabrera/golang-restapi/routes"
)

func main() {

	db.DBConnection()

	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	router := mux.NewRouter()
	router.HandleFunc("/", routes.HomeHandler)

	router.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	router.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	router.HandleFunc("/users", routes.PostUsersHandler).Methods("POST")
	router.HandleFunc("/users", routes.DeleteUsersHandler).Methods("DELETE")

	http.ListenAndServe(":3030", router)
}
