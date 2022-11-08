package repo

import (
	"github.com/gocraft/dbr/v2"

	"github.com/truekupo/cluster/listener-srv/lib/db/models"
)

var (
	txFields = []string{
		"tx.id",
		"tx.created_at",
		"tx.direction",
		"tx.hash",
		"tx.block_num",
		"tx.from_address_id",
		"tx.to_address_id",
		"tx.status",
		"tx.amount",
		"tx.fee",
	}
)

// Repos
type TxRepo struct {
	Repository
}

// Constructors
func Tx(session *dbr.Session) *TxRepo {
	return &TxRepo{New(session, "tx", txFields)}
}

func (r *TxRepo) FindById(id int64) (*models.Tx, error) {
	t := models.Tx{}

	count, err := r.Select().Where("tx.id=?", id).Load(&t)
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, dbr.ErrNotFound
	}

	return &t, nil
}

func (r *TxRepo) FindByAddressId(AddrId int64, From int32, Limit int32) ([]models.Tx, error) {
	t := []models.Tx{}

	count, err := r.Select().Where("tx.from_address_id=? OR tx.to_address_id=?", AddrId, AddrId).Offset(uint64(From)).Limit(uint64(Limit)).OrderAsc("tx.created_at").Load(&t)
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, dbr.ErrNotFound
	}

	return t, nil
}

func (r *TxRepo) FindByTxHash(TxHash string) (*models.Tx, error) {
	t := models.Tx{}

	count, err := r.Select().Where("tx.hash=?", TxHash).Load(&t)
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, dbr.ErrNotFound
	}

	return &t, nil
}

func (r *TxRepo) FindByAddressIds(AddrIds []int64, From int32, Limit int32) ([]models.Tx, error) {
	t := []models.Tx{}

	count, err := r.Select().Where("tx.from_address_id IN ? OR tx.to_address_id IN ?", AddrIds, AddrIds).Offset(uint64(From)).Limit(uint64(Limit)).OrderAsc("tx.created_at").Load(&t)
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, dbr.ErrNotFound
	}

	return t, nil
}

func (r *TxRepo) FindByStatus(Status string) ([]models.Tx, error) {
	t := []models.Tx{}

	count, err := r.Select().Where("tx.status=?", Status).OrderDesc("tx.created_at").Load(&t)
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, dbr.ErrNotFound
	}

	return t, nil

}

func (r *TxRepo) InsertOrUpdate(t *models.Tx) (*models.Tx, error) {
	tt, err := r.Insert(t)
	if err == nil {
		return tt, nil
	}

	return r.Update(t)
}

func (r *TxRepo) Insert(t *models.Tx) (*models.Tx, error) {
	return t, r.InsertBuilder().Columns("created_at", "direction", "hash", "block_num", "from_address_id", "to_address_id", "status", "amount", "fee").Record(t).Returning("id").Load(&t.Id)
}

func (r *TxRepo) Update(t *models.Tx) (*models.Tx, error) {
	_, err := r.UpdateBuilder().Where("hash=?", t.Hash).Set("created_at", t.CreatedAt).Set("direction", t.Direction).Set("block_num", t.BlockNum).Set("status", t.Status).Set("amount", t.Amount).
		Set("fee", t.Fee).Exec()

	return t, err
}

func (r *TxRepo) UpdateStatus(id int64, status string) error {
	return r.update(id, "status", status)
}

func (r *TxRepo) update(id int64, field string, value interface{}) error {
	result, err := r.UpdateBuilder().Where("id=?", id).Set(field, value).Exec()
	if err != nil {
		return err
	}

	count, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if count == 0 {
		return dbr.ErrNotFound
	}

	return nil
}
