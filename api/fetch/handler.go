package fetch

import "net/http"

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":

	default:
		http.Error(w, "404 not found.", http.StatusNotFound)
	}
}

func (h *Handler) FetchEntriesBy(r request) {

}
