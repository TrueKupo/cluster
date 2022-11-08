package repo

import (
	"fmt"

	"github.com/gocraft/dbr/v2"

	"github.com/truekupo/cluster/listener-srv/lib/db/models"
)

var (
	addressFields = []string{
		"address.id",
		"address.active",
		"address.account_id",
		"address.addr",
		"address.request_history",
	}
)

// Repos
type AddressRepo struct {
	Repository
}

// Constructors
func Address(session *dbr.Session) *AddressRepo {
	return &AddressRepo{New(session, "address", addressFields)}
}

func (r *AddressRepo) FindById(id int64) (*models.Address, error) {
	a := models.Address{}

	count, err := r.Select().Where("address.id=?", id).Load(&a)
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, fmt.Errorf("Not found")
	}

	return &a, nil
}

func (r *AddressRepo) AddressById(id int64) string {
	a, err := r.FindById(id)
	if err != nil {
		return ""
	}

	return a.Addr
}

func (r *AddressRepo) FindActive() ([]models.Address, error) {
	a := []models.Address{}

	count, err := r.Select().Where("address.active=true").Load(&a)
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, fmt.Errorf("not found")
	}

	return a, nil
}

func (r *AddressRepo) FindIdByAddress(Address string) (int64, error) {
	a := models.Address{}

	count, err := r.Select().Where("address.active=true and address.addr=?", Address).Load(&a)
	if err != nil {
		return 0, err
	}

	if count == 0 {
		return 0, fmt.Errorf("Not found")
	}

	return a.Id, nil
}

func (r *AddressRepo) FindAddressIdsByAccount(AccountId int64) ([]int64, error) {
	a := []models.Address{}

	count, err := r.Select().Where("address.active=true AND address.account_id=?", AccountId).Load(&a)
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, fmt.Errorf("Not found")
	}

	ids := make([]int64, 0, len(a))
	for _, s := range a {
		ids = append(ids, s.Id)
	}

	return ids, nil
}

func (r *AddressRepo) FindIdByAccountAndAddress(AccountId int64, Address string) (int64, error) {
	a := models.Address{}

	count, err := r.Select().Where("address.active=true and address.account_id=? and address.addr=?", AccountId, Address).Load(&a)
	if err != nil {
		return 0, err
	}

	if count == 0 {
		return 0, fmt.Errorf("Not found")
	}

	return a.Id, nil
}

func (r *AddressRepo) SetHistoryRequested(Id int64) error {
	return r.update(Id, "request_history", true)
}

func (r *AddressRepo) GetAddressesWithHistoryNotRequested() ([]models.Address, error) {
	a := []models.Address{}

	count, err := r.Select().Where("address.request_history=false and address.account_id!=0").Load(&a)
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, fmt.Errorf("not found")
	}

	return a, nil
}

func (r *AddressRepo) GetOrCreateAddressId(AccountId int64, Address string, RequestHistory bool) (int64, error) {
	id, err := r.FindIdByAccountAndAddress(AccountId, Address)
	if err == nil {
		return id, nil
	}

	a, err := r.Insert(&models.Address{
		AccountId:      AccountId,
		Addr:           Address,
		RequestHistory: RequestHistory,
	})
	if err != nil {
		return 0, err
	}

	return a.Id, nil
}

func (r *AddressRepo) Insert(a *models.Address) (*models.Address, error) {
	return a, r.InsertBuilder().Columns("account_id", "addr").Record(a).Returning("id").Load(&a.Id)
}

func (r *AddressRepo) UpdateById(a *models.Address) error {
	result, err := r.UpdateBuilder().Where("id=?", a.Id).Set("account_id", a.AccountId).Set("addr", a.Addr).Exec()
	if err != nil {
		return err
	}

	count, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if count == 0 {
		return fmt.Errorf("Id %d not found", a.Id)
	}

	return nil
}

func (r *AddressRepo) update(id int64, field string, value interface{}) error {
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
