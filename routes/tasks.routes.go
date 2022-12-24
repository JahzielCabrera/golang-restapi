package routes

import (
	"net/http"
)

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GET Tasks Handler"))
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GET Task Handler"))
}

func PostTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("POST Task Handler"))
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DELETE Task Handler"))
}
