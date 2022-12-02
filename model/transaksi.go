package model

type Transaksi struct {
	Saldo   float32 `json:"saldo,omitempty"`
	Gram    float32 `json:"gram,omitempty"`
	HargaID int64   `json:"harga_id,omitempty"`
	Type    uint8   `json:"type,omitempty"`
}
