package model

type Rekening struct {
	Norek string  `json:"norek,omitempty"`
	Saldo float32 `json:"saldo,omitempty"`
	TmplCols
}
