package auth

import (
	"database/sql/driver"
)

// Config represents login config as far as the db is concerned
type Config[T any] struct {
	Cfg T
}

// Value implements the TypeValueScanner.Value method.
func (p *Config[T]) Value() (driver.Value, error) {
	return p.Cfg, nil
}
