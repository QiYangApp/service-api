package auth

// Config represents login config as far as the db is concerned
type Config interface {
	ToDB() ([]byte, error)
	FromDB([]byte) error
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
