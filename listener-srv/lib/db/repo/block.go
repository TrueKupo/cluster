package repo

import (
	"fmt"

	"github.com/gocraft/dbr/v2"

	"github.com/truekupo/cluster/listener-srv/lib/db/models"
)

var (
	blockFields = []string{
		"block.id",
		"block.symbol",
		"block.num",
	}
)

// Repos
type BlockRepo struct {
	Repository
}

// Constructors
func Block(session *dbr.Session) *BlockRepo {
	return &BlockRepo{New(session, "block", blockFields)}
}

func (r *BlockRepo) FindBySymbol(symbol string) (*models.Block, error) {
	b := models.Block{}

	count, err := r.Select().Where("block.symbol=?", symbol).Load(&b)
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, fmt.Errorf("Not found")
	}

	return &b, nil
}

func (r *BlockRepo) BlockNumBySymbol(symbol string) (uint64, error) {
	b, err := r.FindBySymbol(symbol)
	if err != nil {
		return 0, err
	}

	return b.Num, nil
}

func (r *BlockRepo) UpdateBlockNumBySymbol(symbol string, block_num uint64) error {
	return r.UpdateBySymbol(symbol, "num", block_num)
}

func (r *BlockRepo) UpdateBySymbol(symbol string, field string, value interface{}) error {
	result, err := r.UpdateBuilder().Where("symbol=?", symbol).Set(field, value).Exec()
	if err != nil {
		return err
	}

	count, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if count == 0 {
		return fmt.Errorf("Symbol %s not found", symbol)
	}

	return nil
}

func (r *BlockRepo) update(id int64, field string, value interface{}) error {
	result, err := r.UpdateBuilder().Where("id=?", id).Set(field, value).Exec()
	if err != nil {
		return err
	}

	count, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if count == 0 {
		return fmt.Errorf("Id %d not found", id)
	}

	return nil
}
