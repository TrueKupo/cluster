package models

import ()

type Pair struct {
	Id             int64  `json:"id"`
	Active         bool   `json:"active" db:"active"`
	CoinName       string `json:"coin_symbol" db:"coin_symbol"`
	CurrencySymbol string `json:"currency_symbol" db:"currency_symbol"`
}
