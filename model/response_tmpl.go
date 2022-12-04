package model

import (
	"encoding/json"
	"errors"
)

type LastInsertedID struct {
	ID int64 `json:"id"`
}

type ResponseTmpl struct {
	IsError bool   `json:"error"`
	ReffID  string `json:"reff_id,omitempty"`
	Error   error  `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

type inAndOut struct {
	IsError bool   `json:"error"`
	ReffID  string `json:"reff_id,omitempty"`
	Error   string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func (rt ResponseTmpl) MarshalJSON() ([]byte, error) {
	outJSON := inAndOut{IsError: rt.IsError, ReffID: rt.ReffID}
	if rt.Error != nil {
		outJSON.Error = rt.Error.Error()
	}
	if rt.Data != nil {
		outJSON.Data = rt.Data
	}
	return json.Marshal(outJSON)
}

func (rt *ResponseTmpl) UnmarshalJSON(jsonData []byte) error {
	var target inAndOut
	err := json.Unmarshal(jsonData, &target)
	if err != nil {
		return err
	}
	rt.IsError = target.IsError
	rt.ReffID = target.ReffID
	if target.Error != "" {
		rt.Error = errors.New(target.Error)
	}
	if target.Data != nil {
		rt.Data = target.Data
	}
	return nil
}

func NewResponseTmplJSON(isError bool, reffID string, err error, data any) ([]byte, error) {
	respTmpl := ResponseTmpl{IsError: isError, ReffID: reffID, Error: err}
	if data != nil {
		respTmpl.Data = data
	}

	respTmplJSON, err := json.Marshal(respTmpl)
	if err != nil {
		return nil, err
	}
	return respTmplJSON, nil
}
