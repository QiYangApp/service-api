package token

import (
	"crypto/rsa"
	"errors"
	"github.com/imdario/mergo"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

type GeneratorMethodInterface interface {
	init() GeneratorMethodInterface
	generate(claims *Claims) (string, error)
	parser(token string) (*jwt.Token, error)
	refresh(tokenString string) (string, error)
}

type Generator struct {
	ExpiresTime   int
	PrivateKey    string
	PublicKey     string
	PrivateKeyRsa *rsa.PrivateKey
	PublicKeyRsa  *rsa.PublicKey
	Claims        *Claims
	Issuer        string
	Subject       string
	Audience      jwt.ClaimStrings
	ID            string
	Context       interface{}
}

type Claims struct {
	Context interface{} `json:"context"`
	jwt.RegisteredClaims
}

type TokenGenerator struct {
	GeneratorMethodInterface
	*Generator
}

func (t *TokenGenerator) formantKey() {
	var err error
	t.PrivateKeyRsa, err = jwt.ParseRSAPrivateKeyFromPEM([]byte(t.PrivateKey))
	if err == nil {
		zap.S().Fatalf("parse rsa private key failed: %v", err)
	}

	t.PublicKeyRsa, err = jwt.ParseRSAPublicKeyFromPEM([]byte(t.PublicKey))
	if err != nil {
		panic(err)
	}
}

func (t *TokenGenerator) init() GeneratorMethodInterface {
	t.formantKey()

	t.Claims = &Claims{
		t.Context,
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(t.ExpiresTime) * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    t.Issuer,
			Subject:   t.Subject,
			ID:        t.ID,
			Audience:  t.Audience,
		},
	}
	return t
}

func (t TokenGenerator) generate(claims *Claims) (string, error) {
	err := mergo.Merge(t.Claims, claims)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS512, t.Claims)
	return token.SignedString(t.PrivateKey)
}
func (t TokenGenerator) parser(tokenString string) (*jwt.Token, error) {
	// 解析 JWT
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return t.PublicKeyRsa, nil
	})
	if err != nil {
		return nil, err
	}

	// 验证 JWT 签名
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return token, nil
}
func (t TokenGenerator) refresh(tokenString string) (string, error) {
	token, err := t.parser(tokenString)
	if err != nil {
		return "", err
	}

	// 修改 JWT 的有效期。
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return "", errors.New("invalid token claims")
	}

	claims.ExpiresAt = jwt.NewNumericDate(
		time.Now().Add(time.Duration(t.ExpiresTime) * time.Minute),
	)

	// 创建新的 JWT。
	newToken := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)
	newTokenString, err := newToken.SignedString(t.PrivateKeyRsa)
	if err != nil {
		return "", err
	}

	return newTokenString, nil
}
