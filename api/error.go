package api

import (
	"encoding/json"
	"getir-case-study/pkg/utils"
	"log"
	"net/http"
)

func SendError(w http.ResponseWriter, e error) {
	code := 999
	if mappedCode, exists := utils.Errors[e]; exists {
		code = mappedCode
	}

	resJson, err := json.Marshal(&map[string]interface{}{
		"code": code,
		"msg":  e.Error(),
	})
	if err != nil {
		log.Fatal(err)
	}

	w.Write(resJson)
}
