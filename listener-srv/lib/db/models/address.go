package models

import ()

type Address struct {
	Id             int64  `json:"id"`
	Active         bool   `json:"active" db:"active"`
	AccountId      int64  `json:"account_id" db:"account_id"`
	Addr           string `json:"addr" db:"addr"`
	RequestHistory bool   `json:"request_history" db:"request_history"`
}
