package fetch

import (
	"encoding/json"
	"getir-case-study/pkg/model"
	"getir-case-study/pkg/utils"
	"net/http"
)

type response struct {
	Code    int      `json:"code"`
	Message string   `json:"msg"`
	Records []record `json:"records"`
}

type record struct {
	Key        string `json:"key"`
	CreatedAt  string `json:"createdAt"`
	TotalCount int    `json:"totalCount"`
}

func (h *Handler) sendResponse(w http.ResponseWriter, records []model.Record) {
	var resRecords []record
	for i := range records {
		r := records[i]

		resRecords = append(resRecords, record{
			Key:        r.Key,
			CreatedAt:  r.CreatedAt.Format(utils.DateTimeFormat),
			TotalCount: r.TotalCount,
		})
	}

	if resRecords == nil {
		resRecords = []record{}
	}

	resJson, err := json.Marshal(&response{
		Code:    0,
		Message: "success",
		Records: resRecords,
	})
	if err != nil {
		h.sendError(w, err)
	}

	w.Write(resJson)
}
