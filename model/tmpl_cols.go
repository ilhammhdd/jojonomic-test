package model

import "time"

type TmplCols struct {
	ID        int64      `json:"id"`
	CreatedAt *time.Time `json:"created_at"`
}
