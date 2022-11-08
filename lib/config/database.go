package config

import (
	"context"
	"fmt"

	env "github.com/sethvargo/go-envconfig"
)

type Database struct {
	Driver       string `yaml:"driver"`
	Host         string `yaml:"host" env:"DB_HOST,overwrite,default=localhost"`
	Port         int    `yaml:"port" env:"DB_PORT,overwrite,default=5432"`
	DBName       string `yaml:"dbname" env:"DB_NAME,overwrite"`
	User         string `yaml:"user" env:"DB_USER,overwrite"`
	Password     string `yaml:"password" env:"DB_PASSWORD,overwrite"`
	SSLMode      string `yaml:"sslmode"`
	MaxConns     int    `yaml:"max_conns"`
	MaxIdleConns int    `yaml:"max_idle_conns"`
}

func (d *Database) EnvPatch() error {
	return env.Process(context.Background(), d)
}

func (d *Database) Validate() error {
	if d.Driver == "" {
		return fmt.Errorf("Config:Database unknown database.driver")
	}

	if d.DBName == "" {
		return fmt.Errorf("Config:Database unknown database.dbname")
	}

	if d.User == "" {
		return fmt.Errorf("Config:Database unknown database.user")
	}

	if d.Password == "" {
		return fmt.Errorf("Config:Database unknown database.password")
	}

	if d.SSLMode == "" {
		d.SSLMode = "disable"
	}

	if d.MaxConns < 0 {
		d.MaxConns = 0
	}

	if d.MaxIdleConns < 0 {
		d.MaxIdleConns = 0
	}

	if d.MaxIdleConns == 0 {
		d.MaxIdleConns = 2
	}

	if d.Host == "" {
		d.Host = "localhost"
	}

	if d.Port == 0 {
		d.Port = 5432
	}

	return nil
}

func (d *Database) DBDriver() string {
	return d.Driver
}

func (d *Database) DBConnString() string {
	return fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s host=%s port=%d",
		d.User,
		d.Password,
		d.DBName,
		d.SSLMode,
		d.Host,
		d.Port,
	)
}

func (d *Database) DBMaxConns() int {
	return d.MaxConns
}

func (d *Database) DBMaxIdleConns() int {
	return d.MaxIdleConns
}
