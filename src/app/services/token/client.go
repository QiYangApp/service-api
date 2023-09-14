package token

import (
	"github.com/golang-jwt/jwt/v5"
	"service-api/src/core/config"
	"sync"
)

var Instance *Token[GeneratorMethodInterface, CacheMethodInterface]
var once = sync.Once{}

func NewTokenService() *Token[GeneratorMethodInterface, CacheMethodInterface] {

	once.Do(func() {
		tokenConfig := config.Instance.GetToken()

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
	})

	return Instance
}
