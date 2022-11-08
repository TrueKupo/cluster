package middleware

import (
	"github.com/gocraft/dbr/v2"
	_ "github.com/lib/pq"
)

// Database Session Distributor
type DSD struct {
	Conn *dbr.Connection
}

type DbConfiguration interface {
	DBDriver() string
	DBConnString() string
	DBMaxConns() int
	DBMaxIdleConns() int
}

func NewDSD(config DbConfiguration) (*DSD, error) {
	c, err := dbr.Open(config.DBDriver(), config.DBConnString(), nil)
	if err != nil {
		return nil, err
	}

	c.SetMaxIdleConns(config.DBMaxConns())
	c.SetMaxOpenConns(config.DBMaxIdleConns())
	err = c.Ping()
	if err != nil {
		return nil, err
	}

	return &DSD{Conn: c}, nil
}
