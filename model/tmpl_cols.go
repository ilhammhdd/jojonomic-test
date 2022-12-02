package model

import "time"

type TmplCols struct {
	ID        int64      `json:"id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}
