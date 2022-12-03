package model

type TransaksiType uint

const (
	TransaksiType_TOPUP TransaksiType = iota
	TransaksiType_BUYBACK
	TransaksiType_TOPUP_ROLLBACK
	TransaksiType_BUYBACK_ROLLBACK
)

type Transaksi struct {
	Saldo      float32       `json:"saldo,omitempty"`
	Gram       float64       `json:"gram,omitempty"`
	Type       TransaksiType `json:"type"`
	HargaID    int64         `json:"harga_id,omitempty"`
	RekeningID int64         `json:"rekening_id,omitempty"`
}

type TransaksiWithRel struct {
	Transaksi *Transaksi `json:"transaksi"`
	Harga     *Harga     `json:"harga"`
	Rekening  *Rekening  `json:"rekening"`
}
