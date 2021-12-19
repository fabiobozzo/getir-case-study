package fetch

import (
	"encoding/json"
	"getir-case-study/pkg/utils"
	"log"
	"net/http"
)

func (h *Handler) sendError(w http.ResponseWriter, e error) {
	code := 999
	if mappedCode, exists := utils.Errors[e]; exists {
		code = mappedCode
	}

	resJson, err := json.Marshal(&response{
		Code:    code,
		Message: e.Error(),
	})
	if err != nil {
		log.Fatal(err)
	}

	w.Write(resJson)
}
