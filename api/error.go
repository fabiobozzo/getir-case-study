package api

import (
	"encoding/json"
	"getir-case-study/pkg/utils"
	"log"
	"net/http"
)

func SendError(w http.ResponseWriter, e error) {
	code := 999
	if mappedCode, exists := utils.ErrorCodeMap[e]; exists {
		code = mappedCode
	}

	status := 200
	if mappedStatus, exists := utils.ErrorStatusMap[e]; exists {
		status = mappedStatus
	}

	resJson, err := json.Marshal(&map[string]interface{}{
		"code": code,
		"msg":  e.Error(),
	})
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(status)
	w.Write(resJson)
}
