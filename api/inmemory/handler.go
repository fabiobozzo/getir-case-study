package inmemory

import (
	"encoding/json"
	"getir-case-study/api"
	"getir-case-study/pkg/kv"
	"getir-case-study/pkg/utils"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

type Handler struct {
	storage kv.Storage
	logger  *logrus.Logger
}

func NewHandler(logger *logrus.Logger, storage kv.Storage) *Handler {
	return &Handler{
		logger:  logger,
		storage: storage,
	}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		h.storagePut(w, r)
	case "GET":
		h.storageGet(w, r)
	default:

	}
}

// storagePut godoc
// @Summary Insert a key-value pair in the in-mem storage
// @Description Only non-empty values are accepted
// @Tags In-Memory
// @ID in-memory-post
// @Accept json
// @Produce json
// @Param payload body request true "Key-Value Pair"
// @Success 200 {string} string "Ok"
// @Success 400 {string} string "Bad Request"
// @Success 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /in-memory [post]
func (h *Handler) storagePut(w http.ResponseWriter, r *http.Request) {
	if err := api.RequireJson(r); err != nil {
		api.SendError(w, err)

		return
	}
	api.SetJson(w)

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.logger.WithError(err).Error("failed to read request body")
		api.SendError(w, utils.ErrInvalidInput)

		return
	}

	var payload request
	if err := json.Unmarshal(body, &payload); err != nil {
		h.logger.WithError(err).Error("failed to unmarshal json to payload struct")
		api.SendError(w, utils.ErrInvalidInput)

		return
	}

	if err = h.storage.Put(payload.Key, payload.Value); err != nil {
		h.logger.WithError(err).Error("failed to put value in kv storage")
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

// storageGet godoc
// @Summary Find/Read a key-value pair from the in-mem storage
// @Tags In-Memory
// @ID in-memory-get
// @Accept plain
// @Produce json
// @Param key query string true "key of the key-value pair"
// @Success 200 {string} string "Ok"
// @Success 400 {string} string "Bad Request"
// @Success 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /in-memory [get]
func (h *Handler) storageGet(w http.ResponseWriter, r *http.Request) {
	api.SetJson(w)

	key := r.URL.Query().Get("key")
	if len(strings.TrimSpace(key)) == 0 {
		api.SendError(w, utils.ErrInvalidInput)

		return
	}

	value, err := h.storage.Get(key)
	if err != nil {
		h.logger.WithError(err).Error("failed to read value from kv storage")
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
