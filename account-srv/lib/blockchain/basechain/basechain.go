package basechain

import (
	gc "github.com/truekupo/cluster/lib/config"
)

type BaseChain struct {
	conf *gc.Chain
}

func New(conf *gc.Chain) BaseChain {
	return BaseChain{
		conf: conf,
	}
}

func (b *BaseChain) Symbol() string {
	return b.conf.Symbol
}

func (b *BaseChain) Name() string {
	return b.conf.Name
}

func (b *BaseChain) Enabled() bool {
	return b.conf.Enable
}

func (b *BaseChain) Network() string {
	return b.conf.Net
}
