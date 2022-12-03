package model

type TopupRequest struct {
	Gram  float64 `json:"gram"`
	Harga int64   `json:"harga"`
	NoRek string  `json:"norek"`
}
