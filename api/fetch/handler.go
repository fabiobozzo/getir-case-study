package fetch

import (
	"encoding/json"
	"getir-case-study/api"
	"getir-case-study/pkg/db"
	"getir-case-study/pkg/utils"
	"io/ioutil"
	"net/http"
)

type Handler struct {
	dbReader db.Reader
}

func NewHandler(dbReader db.Reader) *Handler {
	return &Handler{
		dbReader: dbReader,
	}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		h.handlePost(w, r)
	default:
		utils.NotFound(w)
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

	valid, startTime, endTime := payload.Validate()
	if !valid {
		api.SendError(w, utils.ErrInvalidInput)

		return
	}

	records, err := h.dbReader.RecordsByDateAndCountRange(
		r.Context(),
		db.DateRange{
			Start: startTime,
			End:   endTime,
		},
		db.CountRange{
			Min: payload.MinCount,
			Max: payload.MaxCount,
		},
	)
	if err != nil {
		api.SendError(w, utils.ErrStorageError)

		return
	}

	sendResponse(w, records)
}
