package model

type Harga struct {
	AdminID string `json:"admin_id,omitempty"`
	Topup   int64  `json:"topup,omitempty"`
	Buyback int64  `json:"buyback,omitempty"`
	TmplCols
}
