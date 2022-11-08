package models

import ()

type Tx struct {
	Id         int64  `json:"id"`
	CreatedAt  uint64 `json:"created_at" db:"created_at"`
	Direction  string `json:"direction" db:"direction"`
	BlockNum   uint64 `json:"block_num" db:"block_num"`
	Hash       string `json:"hash" db:"hash"`
	FromAddrId int64  `json:"from_address_id" db:"from_address_id"`
	ToAddrId   int64  `json:"to_address_id" db:"to_address_id"`
	Status     string `json:"status" db:"status"`
	Amount     string `json:"amount" db:"amount"`
	Fee        string `json:"fee" db:"fee"`
}
