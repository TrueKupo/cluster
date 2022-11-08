package models

import ()

type Block struct {
	Id     int64  `json:"id"`
	Symbol string `json:"symbol" db:"symbol"`
	Num    uint64 `json:"num" db:"num"`
}
