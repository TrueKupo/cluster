package repo

import (
	"fmt"

	"github.com/gocraft/dbr/v2"

	"github.com/truekupo/cluster/coinsrate-srv/lib/db/models"
)

var (
	coinFields = []string{
		"coin.id",
		"coin.symbol",
		"coin.name",
	}
)

// Repos
type CoinRepo struct {
	Repository
}

// Constructors
func Coin(session *dbr.Session) *CoinRepo {
	return &CoinRepo{New(session, "coin", coinFields)}
}

// Methods
func (r *CoinRepo) FindById(id int64) (*models.Coin, error) {
	c := models.Coin{}

	count, err := r.Select().Where("coin.id=?", id).Load(&c)
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, fmt.Errorf("Not found")
	}

	return &c, nil
}

func (r *CoinRepo) Insert(c *models.Coin) (*models.Coin, error) {
	return c, r.InsertBuilder().Columns("symbol", "name").Record(c).Returning("id").Load(&c.Id)
}
