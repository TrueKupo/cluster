package models

import ()

type Kline struct {
	Id        int64  `json:"id"`
	PairId    int64  `json:"pair_id" db:"pair_id"`
	OpenTime  int64  `json:"open_time" db:"open_time"`
	CloseTime int64  `json:"close_time" db:"close_time"`
	Open      string `json:"open" db:"open"`
	Close     string `json:"close" db:"close"`
	High      string `json:"high" db:"high"`
	Low       string `json:"low" db:"low"`
}
