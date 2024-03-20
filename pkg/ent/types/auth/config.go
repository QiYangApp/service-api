package auth

import (
	"database/sql/driver"
)

type Cfg interface {
}

// Config represents login config as far as the db is concerned
type Config struct {
	Cfg Cfg
}

// Value implements the TypeValueScanner.Value method.
func (p *Config) Value() (driver.Value, error) {
	return p.Cfg, nil
}
