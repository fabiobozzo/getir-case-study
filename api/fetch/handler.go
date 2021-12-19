package fetch

import (
	"encoding/json"
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
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			h.sendError(w, utils.ErrInvalidInput)

			return
		}

		var payload request
		if err := json.Unmarshal(body, &payload); err != nil {
			h.sendError(w, utils.ErrInvalidInput)

			return
		}

		valid, startTime, endTime := payload.Validate()
		if !valid {
			h.sendError(w, utils.ErrInvalidInput)

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
			h.sendError(w, utils.ErrDatabaseError)

			return
		}

		h.sendResponse(w, records)
	default:
		utils.NotFound(w)
	}
}
