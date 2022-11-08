package blockchain

import (
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

type Blockachain interface {
	Symbol() string
	Name() string
	Enabled() bool
	Network() string
	SeedDeriveToPublicKeyHex([]byte, string) (string, error)
	SeedDeriveToPrivateKeyHex([]byte, string) (string, error)
	SeedDeriveToAccountData([]byte, string) (string, string, string, error)
}

func NewMnemonic(entropy_size int, lang_id int32) (string, error) {
	entropy, err := bip39.NewEntropy(entropy_size)
	if err != nil {
		return "", err
	}

	bip39.SetWordList(WordList(lang_id))

	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return "", err
	}

	return mnemonic, nil
}

func GetSeed(mnemonic string, password string) ([]byte, error) {
	return bip39.NewSeedWithErrorChecking(mnemonic, password)
}

func GetMasterKey(mnemonic string, password string) ([]byte, error) {
	seed := bip39.NewSeed(mnemonic, password)

	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		return []byte{}, err
	}

	return masterKey.Serialize()
}
