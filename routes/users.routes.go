package routes

import "net/http"

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GET  User Handler"))
}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get Users Handler"))
}

func PostUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("POST Users Handler"))
}

func DeleteUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DELETE Users Handler"))
}
