package utils

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ilhammhdd/jojonomic_test/model"
)

func IsNotNilAndWriteErrResp(statusCode int, w *http.ResponseWriter, err error) bool {
	if err != nil {
		log.Println(err)
		(*w).WriteHeader(statusCode)
		respBodyJSON, err := json.Marshal(model.ResponseTmpl{IsError: true, Error: err})
		if err != nil {
			(*w).Write([]byte(err.Error()))
			return true
		}
		(*w).Write(respBodyJSON)
		return true
	}
	return false
}
