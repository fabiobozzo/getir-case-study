package api

import (
	"getir-case-study/pkg/utils"
	"net/http"
)

func RequireJson(r *http.Request) error {
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" && contentType != "text/json" {
		return utils.ErrInvalidHeaders
	}

	return nil
}

func SetJson(w http.ResponseWriter) {
	w.Header().Set("Content-type", "application/json")
}
