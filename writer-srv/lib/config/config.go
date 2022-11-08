package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"

	cc "github.com/truekupo/cluster/lib/config"
)

type Config struct {
	Global
	System
	cc.Cluster
	cc.Database
	Settings
}

func Parse(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return loadAndValidate(data)
}

func EnvPatch(config *Config) error {
	if err := config.Global.EnvPatch(); err != nil {
		return err
	}

	if err := config.System.EnvPatch(); err != nil {
		return err
	}

	if err := config.Database.EnvPatch(); err != nil {
		return err
	}

	if err := config.Settings.EnvPatch(); err != nil {
		return err
	}

	if err := config.Cluster.EnvPatch(); err != nil {
		return err
	}

	return nil
}

func Validate(config *Config) error {
	if err := config.Global.Validate(); err != nil {
		return err
	}

	if err := config.System.Validate(); err != nil {
		return err
	}

	if err := config.Database.Validate(); err != nil {
		return err
	}

	if err := config.Settings.Validate(); err != nil {
		return err
	}

	if err := config.Cluster.Validate(); err != nil {
		return err
	}

	return nil
}

func loadYaml(data []byte) (*Config, error) {
	config := &Config{}
	if err := yaml.Unmarshal(data, config); err != nil {
		return nil, err
	}
	return config, nil
}

func loadAndValidate(data []byte) (*Config, error) {
	config, err := loadYaml(data)
	if err != nil {
		return nil, err
	}

	if err := EnvPatch(config); err != nil {
		return nil, err
	}

	if err := Validate(config); err != nil {
		return nil, err
	}

	return config, nil
}
