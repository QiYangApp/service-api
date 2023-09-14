package token

import (
	"service-api/src/core/logger"
)

type Token[G GeneratorMethodInterface, C CacheMethodInterface] struct {
	generator    G
	cacheStorage C
}

func (t *Token[G, C]) SetGenerator(generator G) *Token[G, C] {
	t.generator = generator

	return t
}

func (t *Token[G, C]) SetCacheStorage(cache C) *Token[G, C] {
	t.cacheStorage = cache

	return t
}

func (t *Token[G, C]) Exists(token string) bool {
	return t.cacheStorage.exists(token)
}

func (t *Token[G, C]) Generate(claims *Claims) (string, error) {
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

func (t *Token[G, C]) Remove(token string) bool {
	return t.cacheStorage.remove(token)
}

func (t *Token[G, C]) Parser(token string) (*Claims, error) {
	token, err := t.cacheStorage.get(token)
	if err != nil {
		return new(Claims), err
	}

	claims, err := t.generator.parserClaim(token)
	if err != nil {
		return new(Claims), err
	}

	return claims, err
}

func (t *Token[G, C]) Refresh(token string) bool {
	if t.cacheStorage.exists(token) == false {
		return false
	}

	if t.generator.exists(token) == false {
		return false
	}

	newTokenString, err := t.generator.refresh(token)
	if err != nil {
		logger.D().Warnf("token refresh store failed key %s", token)
		return false
	}

	if t.cacheStorage.refresh(token, newTokenString) == false {
		logger.D().Warnf("token refresh cache store failed key %s", token)
		return false
	}

	return true
}
