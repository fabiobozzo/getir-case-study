package utils

import "net/http"

func NotFound(w http.ResponseWriter) {
	http.Error(w, "Not found.", http.StatusNotFound)
}

func BadRequest(w http.ResponseWriter) {
	http.Error(w, "Bad Request.", http.StatusBadRequest)
}

func InternalError(w http.ResponseWriter) {
	http.Error(w, "Internal Server Error.", http.StatusInternalServerError)
}
