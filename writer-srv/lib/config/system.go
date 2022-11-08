package config

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"

	env "github.com/sethvargo/go-envconfig"
)

var (
	logLevel logrus.Level
)

type System struct {
	LogLevel  string `yaml:"log_level" env:"LOG_LEVEL"`
	LogFile   string `yaml:"log_file"`
	LogFormat string `yaml:"log_format" env:"LOG_FORMAT"`
}

func (s *System) EnvPatch() error {
	return env.Process(context.Background(), s)
}

func (s *System) Validate() error {
	if s.LogFile == "" {
		return fmt.Errorf("Config:System unknown system.log_file")
	}

	// Trace, Debug, Info, Warn, Error, Fatal or Panic
	if s.LogLevel == "" {
		s.LogLevel = "error"
	}

	// text, json
	if s.LogFormat != "text" {
		s.LogFormat = "json"
	}

	ll, err := logrus.ParseLevel(s.LogLevel)
	if err != nil {
		ll = logrus.ErrorLevel
	}

	logLevel = ll

	return nil
}

func (s *System) Fields() logrus.Fields {
	return logrus.Fields{
		"log_file":  s.LogFile,
		"log_level": s.LogLevel,
	}
}
