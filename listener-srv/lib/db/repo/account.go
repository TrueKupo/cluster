package repo

import (
	"fmt"

	"github.com/gocraft/dbr/v2"

	"github.com/truekupo/cluster/listener-srv/lib/db/models"
)

var (
	accountFields = []string{
		"account.id",
		"account.active",
		"account.uuid",
	}
)

// Repos
type AccountRepo struct {
	Repository
}

// Constructors
func Account(session *dbr.Session) *AccountRepo {
	return &AccountRepo{New(session, "account", accountFields)}
}

func (r *AccountRepo) FindIdByUuid(uuid string) (int64, error) {
	w, err := r.FindByUuid(uuid)
	if err != nil {
		return 0, err
	}

	return w.Id, nil
}

func (r *AccountRepo) GetOrCreateAccountId(uuid string) (int64, error) {
	id, err := r.FindIdByUuid(uuid)
	if err == nil {
		return id, nil
	}

	w := &models.Account{
		Uuid: uuid,
	}

	w, err = r.Insert(w)
	if err != nil {
		return 0, err
	}

	return w.Id, nil
}

func (r *AccountRepo) FindById(id int64) (*models.Account, error) {
	w := models.Account{}

	count, err := r.Select().Where("account.id=?", id).Load(&w)
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, fmt.Errorf("Not found")
	}

	return &w, nil
}

func (r *AccountRepo) FindByUuid(uuid string) (*models.Account, error) {
	w := models.Account{}

	count, err := r.Select().Where("account.uuid=?", uuid).Load(&w)
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, fmt.Errorf("Not found")
	}

	return &w, nil
}

func (r *AccountRepo) Insert(w *models.Account) (*models.Account, error) {
	return w, r.InsertBuilder().Columns("uuid").Record(w).Returning("id").Load(&w.Id)
}

func (r *AccountRepo) update(id int64, field string, value interface{}) error {
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
