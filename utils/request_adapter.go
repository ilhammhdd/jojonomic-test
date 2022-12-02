package utils

import (
	"encoding/json"
	"net/http"
)

func ReadRequestBodyJSON[T interface{}](r *http.Request) (*T, error) {
	bodyData := make([]byte, r.ContentLength)
	r.Body.Read(bodyData)
	defer r.Body.Close()

	var result T
	err := json.Unmarshal(bodyData, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
