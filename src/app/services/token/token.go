package token

import (
	"github.com/golang-jwt/jwt/v5"
	"service-api/src/app/helpers"
	"service-api/src/core/system"
)

var Instance *Token[GeneratorMethodInterface, CacheMethodInterface]

func init() {
	NewTokenService()
}

type Token[G GeneratorMethodInterface, C CacheMethodInterface] struct {
	generator    G
	cacheStorage C
}

func (t *Token[G, C]) setGenerator(generator G) *Token[G, C] {
	t.generator = generator

	return t
}

func (t *Token[G, C]) setCacheStorage(cache C) *Token[G, C] {
	t.cacheStorage = cache

	return t
}

func (t *Token[G, C]) exists(token string) bool {
	return t.cacheStorage.exists(token)
}

func (t *Token[G, C]) generate(claims *Claims) (string, error) {
	token, err := t.generator.generate(claims)
	if err != nil {
		return "", err
	}

	token, err = t.cacheStorage.store(token)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (t *Token[G, C]) remove(token string) bool {
	return t.cacheStorage.remove(token)
}

func (t *Token[G, C]) parser(token string) (Claims, error) {
	return t.cacheStorage.parser(token)
}

func NewTokenService() *Token[GeneratorMethodInterface, CacheMethodInterface] {
	if helpers.IsEmpty(Instance) == true {
		tokenConfig := system.ConfigInstance.GetToken()

		Instance = &Token[GeneratorMethodInterface, CacheMethodInterface]{
			generator: &TokenGenerator{
				Generator: &Generator{
					ExpiresTime: tokenConfig.ExpiresTime,
					Subject:     "default",
					Audience:    jwt.ClaimStrings{"default", "system"},
					PrivateKey:  tokenConfig.PrivateKey,
					PublicKey:   tokenConfig.PublicKey,
				},
			},
			cacheStorage: &CacheRedis{
				Cache: &Cache{
					ExpiresTime: tokenConfig.ExpiresTime,
				},
			},
		}

		Instance.generator.init()
		Instance.cacheStorage.init()
	}

	return Instance
}
