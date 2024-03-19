package auth

import (
	"database/sql/driver"
	"service-api/internal/ent"
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

// SkipVerifiable configurations provide a IsSkipVerify to check if SkipVerify is set
type SkipVerifiable interface {
	IsSkipVerify() bool
}

// HasTLSer configurations provide a HasTLS to check if TLS can be enabled
type HasTLSer interface {
	HasTLS() bool
}

// UseTLSer configurations provide a HasTLS to check if TLS is enabled
type UseTLSer interface {
	UseTLS() bool
}

// SSHKeyProvider configurations provide ProvidesSSHKeys to check if they provide SSHKeys
type SSHKeyProvider interface {
	ProvidesSSHKeys() bool
}

// RegisterableSource configurations provide RegisterSource which needs to be run on creation
type RegisterableSource interface {
	RegisterSource() error
	UnregisterSource() error
}

// SourceSettable configurations can have their authSource set on them
type SourceSettable interface {
	SetAuthSource(*ent.Source)
}
