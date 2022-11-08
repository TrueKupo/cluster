package config

import (
	"context"

	env "github.com/sethvargo/go-envconfig"
)

type Listener struct {
	Symbol string `yaml:"symbol"`
	Addr   string `yaml:"addr" env:"LISTENER_ADDR_PORT,overwrite"`
}

type Writer struct {
	Addr string `yaml:"addr" env:"WRITER_ADDR_PORT,overwrite"`
}

type Cluster struct {
	Listeners []Listener `json:"listener" yaml:"listener"`
	Writer    Writer     `json:"writer" yaml:"writer"`
}

func (l *Listener) EnvPatch() error {
	e := env.PrefixLookuper(l.Symbol+"_", env.OsLookuper())
	return env.ProcessWith(context.Background(), l, e)
}

func (w *Writer) EnvPatch() error {
	return env.Process(context.Background(), w)
}

func (c Cluster) EnvPatch() error {
	for i := 0; i < len(c.Listeners); i++ {
		listener := &c.Listeners[i]
		err := listener.EnvPatch()
		if err != nil {
			return err
		}
	}

	err := c.Writer.EnvPatch()
	if err != nil {
		return err
	}

	return nil
}

func (c Cluster) Validate() error {
	return nil
}
