package listener

import (
	"fmt"
)

type addressData struct {
	address_id int64
	account_id int64
}

type AddressWatcher struct {
	w map[string]addressData
}

func NewAddressWatcher() *AddressWatcher {
	return &AddressWatcher{
		w: map[string]addressData{},
	}
}

func (a *AddressWatcher) Add(Address string, AddressId int64, AccountId int64) error {
	a.w[Address] = addressData{
		address_id: AddressId,
		account_id: AccountId,
	}
	return nil
}

func (a *AddressWatcher) Get(Address string) (int64, int64, error) {
	d, ok := a.w[Address]
	if !ok {
		return 0, 0, fmt.Errorf("Not found")
	}

	return d.address_id, d.account_id, nil
}
