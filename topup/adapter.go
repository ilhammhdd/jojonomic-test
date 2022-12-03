package main

import (
	"encoding/json"

	"github.com/ilhammhdd/jojonomic_test/model"
)

func ConvertToByteVal(t *model.TransaksiWithRel) ([]byte, error) {
	requestJSON, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return requestJSON, nil
}
