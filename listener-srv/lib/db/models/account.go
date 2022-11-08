package models

import ()

type Account struct {
	Id     int64  `json:"id"`
	Active string `json:"active" db:"active"`
	Uuid   string `json:"uuid" db:"uuid"`
}
