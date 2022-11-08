package repo

import (
	"fmt"

	"github.com/gocraft/dbr/v2"

	"github.com/truekupo/cluster/coinsrate-srv/lib/db/models"
)

var (
	klineFields = []string{
		"kline.id",
		"kline.pair_id",
		"kline.open_time",
		"kline.close_time",
		"kline.open",
		"kline.close",
		"kline.high",
		"kline.low",
	}
)

// Repos
type KlineRepo struct {
	Repository
}

// Constructors
func Kline(session *dbr.Session) *KlineRepo {
	return &KlineRepo{New(session, "kline", klineFields)}
}

// Methods
func (r *KlineRepo) FindById(id int64) (*models.Kline, error) {
	k := models.Kline{}

	count, err := r.Select().Where("kline.id=?", id).Load(&k)
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, fmt.Errorf("Not found")
	}

	return &k, nil
}

func (r *KlineRepo) FindLast(PairId int64) (*models.Kline, error) {
	k := models.Kline{}

	count, err := r.Select().Where("kline.pair_id=?", PairId).OrderDesc("kline.close_time").Limit(1).Load(&k)
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, fmt.Errorf("Not found")
	}

	return &k, nil
}

func (r *KlineRepo) Insert(k *models.Kline) (*models.Kline, error) {
	return k, r.InsertBuilder().Columns("pair_id", "open_time", "close_time", "open", "close", "high", "low").Record(k).Returning("id").Load(&k.Id)
}
