package fetch

import (
	"encoding/json"
	"getir-case-study/api"
	"getir-case-study/pkg/db"
	"getir-case-study/pkg/utils"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

type Handler struct {
	logger   *logrus.Logger
	dbReader db.Reader
}

func NewHandler(logger *logrus.Logger, dbReader db.Reader) *Handler {
	return &Handler{
		logger:   logger,
		dbReader: dbReader,
	}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		h.fetch(w, r)
	default:
		utils.NotFound(w)
	}
}

// fetch godoc
// @Summary Fetch records from DB
// @Description Filter records by date and total count range
// @Tags Fetch
// @ID fetch-post
// @Accept json
// @Produce json
// @Param payload body request true "Search Filter"
// @Success 200 {string} string "Ok"
// @Success 400 {string} string "Bad Request"
// @Success 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /fetch [post]
func (h *Handler) fetch(w http.ResponseWriter, r *http.Request) {
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
		h.logger.WithError(err).Error("db error: RecordsByDateAndCountRange")
		api.SendError(w, utils.ErrStorageError)

		return
	}

	sendResponse(w, records)
}
