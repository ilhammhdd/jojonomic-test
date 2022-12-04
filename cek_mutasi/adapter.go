package main

type inRequestBody struct {
	Norek     string `json:"norek"`
	StartDate int64  `json:"start_date"`
	EndDate   int64  `json:"end_date"`
}

type outTransaction struct {
	Date         int64   `json:"date"`
	Type         string  `json:"type"`
	Gram         float64 `json:"gram"`
	HargaTopup   int64   `json:"harga_topup"`
	HargaBuyback int64   `json:"harga_buyback"`
	Saldo        float64 `json:"saldo"`
}
