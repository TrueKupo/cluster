package config

import (
	"context"
	"fmt"

	env "github.com/sethvargo/go-envconfig"
)

type Global struct {
	GrpcListenAddr string `yaml:"grpc_listen_addr" json:"grpc_listen_addr" db:"grpc_listen_addr" env:"GRPC_LISTEN_ADDR_PORT,overwrite"`
}

func (g *Global) EnvPatch() error {
	return env.Process(context.Background(), g)
}

func (g *Global) Validate() error {
	if g.GrpcListenAddr == "" {
		return fmt.Errorf("Config:Global: unknown global.grpc_listen_addr")
	}

	return nil
}
