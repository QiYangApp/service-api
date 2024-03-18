package auth

import (
	"database/sql/driver"
	"encoding/hex"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/schema/field"
	"fmt"
	"strings"
)

type Cfg interface {
}

// Config represents login config as far as the db is concerned
type Config struct {
	context string
}

// Value implements the TypeValueScanner.Value method.
func (p Config) Value(s string) (driver.Value, error) {
	return p.context + ":" + hex.EncodeToString([]byte(s)), nil
}

// ScanValue implements the TypeValueScanner.ScanValue method.
func (Config) ScanValue() field.ValueScanner {
	return &sql.NullString{}
}

// FromValue implements the TypeValueScanner.FromValue method.
func (p Config) FromValue(v driver.Value) (string, error) {
	s, ok := v.(*sql.NullString)
	if !ok {
		return "", fmt.Errorf("unexpected input for FromValue: %T", v)
	}
	if !s.Valid {
		return "", nil
	}
	d, err := hex.DecodeString(strings.TrimPrefix(s.String, p.context+":"))
	if err != nil {
		return "", err
	}
	return string(d), nil
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
