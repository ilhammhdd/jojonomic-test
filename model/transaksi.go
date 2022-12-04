package model

type TransaksiType string

const (
	TransaksiType_TOPUP            string = "topup"
	TransaksiType_BUYBACK          string = "buyback"
	TransaksiType_TOPUP_ROLLBACK   string = "topup_rollback"
	TransaksiType_BUYBACK_ROLLBACK string = "buyback_rollback"
)

type Transaksi struct {
	Saldo      float32 `json:"saldo,omitempty"`
	Gram       float64 `json:"gram,omitempty"`
	Type       string  `json:"type"`
	HargaID    int64   `json:"harga_id,omitempty"`
	RekeningID int64   `json:"rekening_id,omitempty"`
}

type TransaksiWithRel struct {
	Transaksi *Transaksi `json:"transaksi"`
	Harga     *Harga     `json:"harga"`
	Rekening  *Rekening  `json:"rekening"`
}
