package chain

import (
	gc "github.com/truekupo/cluster/lib/config"
)

type Chain struct {
	conf *gc.Chain
}

func New(conf *gc.Chain) Chain {
	return Chain{
		conf: conf,
	}
}

func (b *Chain) Symbol() string {
	return b.conf.Symbol
}

func (b *Chain) Name() string {
	return b.conf.Name
}

func (b *Chain) Enabled() bool {
	return true
}

func (b *Chain) Network() string {
	return b.conf.Net
}

func (b *Chain) Url() string {
	return b.conf.Url
}

func (b *Chain) Secret() string {
	return b.conf.Secret
}
