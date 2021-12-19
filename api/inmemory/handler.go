package inmemory

import (
	"getir-case-study/pkg/kv"
	"net/http"
)

type Handler struct {
	storage kv.Storage
}

func NewHandler(storage kv.Storage) *Handler {
	return &Handler{
		storage: storage,
	}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":

	case "GET":

	default:

	}
}
