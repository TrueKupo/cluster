package repo

import (
	"fmt"

	"github.com/gocraft/dbr/v2"

	"github.com/truekupo/cluster/coinsrate-srv/lib/db/models"
)

var (
	pairFields = []string{
		"pair.id",
		"pair.active",
		"pair.coin_symbol",
		"pair.currency_symbol",
	}
)

// Repos
type PairRepo struct {
	Repository
}

// Constructors
func Pair(session *dbr.Session) *PairRepo {
	return &PairRepo{New(session, "pair", pairFields)}
}

// Methods
func (r *PairRepo) FindById(id int64) (*models.Pair, error) {
	p := models.Pair{}

	count, err := r.Select().Where("pair.id=?", id).Load(&p)
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, fmt.Errorf("Not found")
	}

	return &p, nil
}

func (r *PairRepo) FindActive() ([]models.Pair, error) {
	p := []models.Pair{}

	count, err := r.Select().Where("pair.active=true").Load(&p)
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, fmt.Errorf("Not found")
	}

	return p, nil
}

func (r *PairRepo) FindByNames(Coin string, Currency string) (*models.Pair, error) {
	p := models.Pair{}

	count, err := r.Select().Where("pair.coin_symbol=? and pair.currency_symbol=?", Coin, Currency).Load(&p)
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, fmt.Errorf("Not found")
	}

	return &p, nil
}

func (r *PairRepo) Insert(p *models.Pair) (*models.Pair, error) {
	return p, r.InsertBuilder().Columns("active", "coin_symbol", "currency_symbol").Record(p).Returning("id").Load(&p.Id)
}

func (r *PairRepo) Name(id int64) (string, error) {
	p, err := r.FindById(id)
	if err != nil {
		return "", err
	}

	return p.CoinName + p.CurrencySymbol, nil
}
