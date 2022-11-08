package chain

import (
	"github.com/shopspring/decimal"
)

// Transaction Directions
const (
	TransactionDirectionIn    = "Input"
	TransactionDirectionOut   = "Output"
	TransactionDirectionInOut = "InputOutput"
)

// Transaction Statuses
const (
	TransactionStatusNew        = "New"
	TransactionStatusProcessing = "Processing"
	TransactionStatusFinalized  = "Finalized"
	TransactionStatusFailed     = "Failed"
)

type Transaction struct {
	CreatedAt uint64
	Hash      string
	BlockNum  uint64
	From      string
	To        string
	Direction string
	Status    string
	Amount    decimal.Decimal
	Fee       decimal.Decimal
}
