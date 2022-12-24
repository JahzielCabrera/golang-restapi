package routes

import (
	"net/http"
)

func Tasks(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Tasks"))
}
