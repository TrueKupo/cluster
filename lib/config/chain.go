package config

import (
	"context"

	env "github.com/sethvargo/go-envconfig"
)

type Chain struct {
	Symbol    string `yaml:"symbol" env:"CHAIN_SYMBOL,overwrite"`
	Name      string `yaml:"name" env:"NAME,overwrite"`
	Enable    bool   `yaml:"enable" env:"ENABLE,overwrite"`
	Net       string `yaml:"net" env:"NET,overwrite"`
	Url       string `yaml:"url" env:"URL,overwrite"`
	RequestTm string `yaml:"request_timeout" env:"REQUEST_TIMEOUT,overwrite"`
	Secret    string `yaml:"secret" env:"SECRET,overwrite"`
}

func (c *Chain) EnvPatch() error {
	err := env.Process(context.Background(), c)
	if err != nil {
		return err
	}

	return env.ProcessWith(context.Background(), c, env.PrefixLookuper(c.Symbol+"_", env.OsLookuper()))

}

func (c *Chain) Validate() error {
	return nil
}
