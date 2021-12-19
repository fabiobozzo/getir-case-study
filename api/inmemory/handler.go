package inmemory

import (
	"encoding/json"
	"getir-case-study/api"
	"getir-case-study/pkg/kv"
	"getir-case-study/pkg/utils"
	"io/ioutil"
	"net/http"
	"strings"
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
		h.handlePost(w, r)
	case "GET":
		h.handleGet(w, r)
	default:

	}
}

func (h *Handler) handlePost(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		api.SendError(w, utils.ErrInvalidInput)

		return
	}

	var payload request
	if err := json.Unmarshal(body, &payload); err != nil {
		api.SendError(w, utils.ErrInvalidInput)

		return
	}

	if err = h.storage.Put(payload.Key, payload.Value); err != nil {
		api.SendError(w, utils.ErrStorageError)

		return
	}

	resJson, err := json.Marshal(&responsePost{
		Code:    0,
		Message: "success",
	})
	if err != nil {
		api.SendError(w, utils.ErrInternalError)

		return
	}

	w.Write(resJson)
}

func (h *Handler) handleGet(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if len(strings.TrimSpace(key)) == 0 {
		api.SendError(w, utils.ErrInvalidInput)

		return
	}

	value, err := h.storage.Get(key)
	if err != nil {
		api.SendError(w, utils.ErrKeyNotFound)

		return
	}

	resJson, err := json.Marshal(&responseGet{
		Key:   key,
		Value: value,
	})
	if err != nil {
		api.SendError(w, utils.ErrInternalError)

		return
	}

	w.Write(resJson)
}
