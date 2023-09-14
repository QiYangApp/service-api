package token

import (
	"crypto/rsa"
	"dario.cat/mergo"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"service-api/src/core/helpers"
	"service-api/src/core/logger"
	"time"
)

type GeneratorMethodInterface interface {
	init() GeneratorMethodInterface
	generate(claims *Claims) (string, error)
	parser(token string) (*jwt.Token, error)
	parserClaim(token string) (*Claims, error)
	refresh(tokenString string) (string, error)
	exists(tokenString string) bool
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

	// Load RSA private key
	privKey, err := os.ReadFile(helpers.PathInstance.JoinCurrentRunPath(t.PrivateKey))
	if err != nil {
		logger.D().Fatalf("Failed to load RSA private key: %v", err)
		return
	}

	// Load RSA private key
	pubKey, err := os.ReadFile(helpers.PathInstance.JoinCurrentRunPath(t.PublicKey))
	if err != nil {
		logger.D().Fatalf("Failed to load RSA private key: %v", err)
		return
	}

	t.PrivateKeyRsa, err = jwt.ParseRSAPrivateKeyFromPEM(privKey)
	if err != nil {
		logger.D().Fatalf("parse rsa private key failed: %v", err)
		return
	}

	t.PublicKeyRsa, err = jwt.ParseRSAPublicKeyFromPEM(pubKey)
	if err != nil {
		logger.D().Fatalf("parse rsa public key failed: %v", err)
		return
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

	t.Claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Duration(t.ExpiresTime) * time.Minute))
	t.Claims.IssuedAt = jwt.NewNumericDate(time.Now())
	t.Claims.NotBefore = jwt.NewNumericDate(time.Now())
	token := jwt.NewWithClaims(jwt.SigningMethodRS512, t.Claims)
	return token.SignedString(t.PrivateKeyRsa)
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

func (t TokenGenerator) exists(tokenString string) bool {
	_, err := t.parser(tokenString)
	if err != nil {
		return false
	}

	return true
}

func (t TokenGenerator) parserClaim(token string) (*Claims, error) {
	parser, err := t.parser(token)
	if err != nil {
		return new(Claims), err
	}

	// 修改 JWT 的有效期。
	return parser.Claims.(*Claims), nil
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
