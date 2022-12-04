package model

type TransaksiRequest struct {
	Gram  float64 `json:"gram"`
	Harga int64   `json:"harga"`
	NoRek string  `json:"norek"`
}
