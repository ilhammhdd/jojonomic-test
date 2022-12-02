package model

import (
	"encoding/json"
	"errors"
)

type ResponseTmpl struct {
	IsError bool   `json:"error"`
	ReffID  string `json:"reff_id,omitempty"`
	Error   error  `json:"message,omitempty"`
}

type inAndOut struct {
	IsError bool   `json:"error"`
	ReffID  string `json:"reff_id,omitempty"`
	Error   string `json:"message,omitempty"`
}

func (rt ResponseTmpl) MarshalJSON() ([]byte, error) {
	outJSON := inAndOut{IsError: rt.IsError, ReffID: rt.ReffID}
	if rt.Error != nil {
		outJSON.Error = rt.Error.Error()
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
	rt.Error = errors.New(target.Error)
	return nil
}
