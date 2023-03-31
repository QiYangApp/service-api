package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Token[G GeneratorInterface] struct {
	generator G
}

type GeneratorInterface interface {
	GeneratorMethodInterface
}

type GeneratorMethodInterface interface {
	init() *Generator
	generate() string
	parser() interface{}
	updateAt() bool
}

type Generator struct {
	GeneratorMethodInterface
	ExpiresAt int
	Issuer    string
	Subject   string
	ID        string
	Audience  []string
	jwt.RegisteredClaims
}

type TokenGenerator struct {
	*Generator
}

func (t *TokenGenerator) init() *Generator {
	// Create the claims

	t.RegisteredClaims = jwt.RegisteredClaims{
		// A usual scenario is to set the expiration time relative to the current time
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(12 * time.Minute)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		Issuer:    t.Issuer,
		Subject:   t.Subject,
		ID:        t.ID,
		Audience:  t.Audience,
	}

	return t.Generator
}

func (t *TokenGenerator) generate() string {
	return ""
}
func (t *TokenGenerator) parser() interface{} {
	return new(interface{})
}
func (t *TokenGenerator) updateAt() bool {
	return true
}

func NewTokenService() *Token[GeneratorInterface] {
	return &Token[GeneratorInterface]{
		generator: &TokenGenerator{&Generator{}},
	}
}
