package models

import ()

type Coin struct {
	Id     int64  `json:"id"`
	Symbol string `json:"symbol" db:"symbol"`
	Name   string `json:"name" db:"name"`
}
