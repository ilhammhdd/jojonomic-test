package model

type Rekening struct {
	Norek string  `json:"norek,omitempty"`
	Saldo float64 `json:"saldo,omitempty"`
	TmplCols
}
