package token

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"service-api/src/core/cache"
	"service-api/src/core/logger"
)

type CacheMethodInterface interface {
	init() CacheMethodInterface
	store(token string) (string, error)
	get(token string) (string, error)
	exists(token string) bool
	refresh(token string, newToken string) bool
	remove(token string) bool
}

type Cache struct {
	ExpiresTime int
}

type CacheRedis struct {
	CacheMethodInterface
	*Cache
}

func (c *CacheRedis) init() CacheMethodInterface {

	return c
}

func (c *CacheRedis) store(token string) (string, error) {
	key := c.generate(token)
	if cache.SetNx(key, token, c.ExpiresTime) == false {
		logger.D().Errorf("token cache failed, token: %s", token)
		return "", errors.New("token cache failed")
	}

	return key, nil
}

func (c *CacheRedis) get(token string) (string, error) {
	return cache.Get(token)
}

func (c *CacheRedis) exists(token string) bool {
	return cache.Exists(token)
}
func (c *CacheRedis) refresh(token, newToken string) bool {
	if cache.SetEx(token, newToken, c.ExpiresTime) == false {
		logger.D().Errorf("token cache refresh failed, token: %s", token)
		return false
	}

	return true
}
func (c *CacheRedis) remove(token string) bool {
	if cache.Exists(token) == false {
		return false
	}

	return cache.Del(token)
}

func (c *CacheRedis) generate(token string) string {
	hash := md5.Sum([]byte(token))

	// 将 md5 值转换为字符串
	return hex.EncodeToString(hash[:])
}
