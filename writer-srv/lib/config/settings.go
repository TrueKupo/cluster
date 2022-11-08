package config

import (
	gc "github.com/truekupo/cluster/lib/config"
)

type Settings struct {
	Chains []gc.Chain `yaml:"chains"`
}

func (s *Settings) EnvPatch() error {
	for i := 0; i < len(s.Chains); i++ {
		chain := &s.Chains[i]
		if err := chain.EnvPatch(); err != nil {
			return err
		}
	}

	return nil
}

func (s *Settings) Validate() error {
	return nil
}
