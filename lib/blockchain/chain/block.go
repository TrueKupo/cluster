package chain

import ()

type Block struct {
	Num       uint64
	Code      string
	CreatedAt uint64
	TxList    []Transaction
}

func NewBlock(num uint64, code string, tm uint64) *Block {
	return &Block{
		Num:       num,
		Code:      code,
		CreatedAt: tm,
		TxList:    []Transaction{},
	}
}

func (b *Block) AddTx(tx Transaction) {
	b.TxList = append(b.TxList, tx)
}
