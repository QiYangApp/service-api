package cache

import (
	"fmt"
	"service-api/src/core/config"
	"service-api/src/core/logger"
	"time"

	"github.com/redis/go-redis/v9"
)

type CacheRedisMethod interface {
}

type CacheRedisDrive struct {
	storage *redis.Client
	CacheDrive
}

func (c *CacheRedisDrive) Connect(cfg interface{}, cCfg interface{}) error {
	nCfg := cfg.(config.DatabaseRedis)
	nCCfg := cCfg.(config.CacheRedis)
	c.storage = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", nCfg.Host, nCfg.Port),
		Password: nCfg.Password,  // no password set
		DB:       nCCfg.Database, // use default DB
	})

	return c.storage.Ping(ctx).Err()
}

func (c *CacheRedisDrive) Get(key string) (string, error) {
	return c.storage.Get(ctx, key).Result()
}

func (c *CacheRedisDrive) SetNx(key string, val interface{}, exp int) (bool, error) {
	return c.storage.SetNX(ctx, key, val, time.Duration(exp)*time.Minute).Result()
}

func (c *CacheRedisDrive) SetEx(key string, val interface{}, exp int) (string, error) {
	return c.storage.SetEx(ctx, key, val, time.Duration(exp)*time.Minute).Result()
}

func (c *CacheRedisDrive) Exists(key string) bool {
	state := c.storage.Exists(ctx, key)
	if state.Err() != nil {
		return false
	}

	return state.Val() == 1
}

func (c *CacheRedisDrive) Refresh(key string, exp int) bool {

	if c.Exists(key) == false {
		return false
	}

	valState, err := c.Get(key)
	if err != nil {
		logger.S().Warnf("cache drive refresh fail, get origin data fail, key, %s", err)
		return false
	}

	state, err := c.SetNx(key, valState, exp)
	if err != nil {
		logger.S().Warnf("cache drive refresh fail, set new data fail, key, %s", err)
		return false
	}

	return state
}

func (c *CacheRedisDrive) Del(key string) bool {
	return c.storage.Del(ctx, key).Val() == 1
}
