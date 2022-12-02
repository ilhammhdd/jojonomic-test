package main

type Request struct {
	AdminID      string `json:"admin_id,omitempty"`
	HargaTopup   int64  `json:"harga_topup,omitempty"`
	HargaBuyback int64  `json:"harga_buyback,omitempty"`
}
