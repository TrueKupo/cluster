package solana

import (
	"github.com/btcsuite/btcutil/base58"
	"github.com/portto/solana-go-sdk/types"
	"github.com/stellar/go/exp/crypto/derivation"

	"github.com/truekupo/cluster/account-srv/lib/blockchain/basechain"
	gc "github.com/truekupo/cluster/lib/config"
)

type Solana struct {
	basechain.BaseChain
}

func New(conf *gc.Chain) (*Solana, error) {
	return &Solana{
		basechain.New(conf),
	}, nil
}

func (e *Solana) SeedDeriveToPublicKeyHex(Seed []byte, Path string) (string, error) {
	kk, err := derivation.DeriveForPath(Path, Seed)
	if err != nil {
		return "", err
	}

	wallet, err := types.AccountFromSeed(kk.Key)
	if err != nil {
		return "", err
	}

	return wallet.PublicKey.ToBase58(), nil
}

func (e *Solana) SeedDeriveToPrivateKeyHex(Seed []byte, Path string) (string, error) {
	kk, err := derivation.DeriveForPath(Path, Seed)
	if err != nil {
		return "", err
	}

	wallet, err := types.AccountFromSeed(kk.Key)
	if err != nil {
		return "", err
	}

	return base58.Encode(wallet.PrivateKey), nil
}

func (e *Solana) SeedDeriveToAccountData(Seed []byte, Path string) (string, string, string, error) {
	kk, err := derivation.DeriveForPath(Path, Seed)
	if err != nil {
		return "", "", "", err
	}

	wallet, err := types.AccountFromSeed(kk.Key)
	if err != nil {
		return "", "", "", err
	}

	return base58.Encode(wallet.PrivateKey), wallet.PublicKey.ToBase58(), wallet.PublicKey.ToBase58(), nil
}
