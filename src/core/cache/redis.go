package cache

import (
	"context"
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

func (c *CacheRedisDrive) Connect(ctx context.Context, cfg interface{}, cCfg interface{}) error {
	nCfg := cfg.(config.DatabaseRedis)
	nCCfg := cCfg.(config.CacheRedis)
	c.storage = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", nCfg.Host, nCfg.Port),
		Password: nCfg.Password,  // no password set
		DB:       nCCfg.Database, // use default DB
	})

	_, err := c.storage.Conn().Ping(ctx).Result()
	if err != nil {
		return err
	}

	_, err = c.storage.Conn().Auth(ctx, nCfg.Password).Result()
	if err != nil {
		return err
	}

	return nil
}

func (c *CacheRedisDrive) Get(ctx context.Context, key string) (string, error) {
	return c.storage.Get(ctx, key).Result()
}

func (c *CacheRedisDrive) SetNx(ctx context.Context, key string, val interface{}, exp time.Duration) (bool, error) {
	return c.storage.SetNX(ctx, key, val, exp).Result()
}

func (c *CacheRedisDrive) SetEx(ctx context.Context, key string, val interface{}, exp time.Duration) (string, error) {
	return c.storage.SetEx(ctx, key, val, exp).Result()
}

func (c *CacheRedisDrive) Exists(ctx context.Context, key string) bool {
	state := c.storage.Exists(ctx, key)
	if state.Err() != nil {
		return false
	}

	return state.Val() == 1
}

func (c *CacheRedisDrive) Refresh(ctx context.Context, key string, exp time.Duration) bool {

	if c.Exists(ctx, key) == false {
		return false
	}

	valState, err := c.Get(ctx, key)
	if err != nil {
		logger.S().Warnf("cache drive refresh fail, get origin data fail, key, %s", err)
		return false
	}

	state, err := c.SetNx(ctx, key, valState, exp)
	if err != nil {
		logger.S().Warnf("cache drive refresh fail, set new data fail, key, %s", err)
		return false
	}

	return state
}

func (c *CacheRedisDrive) Del(ctx context.Context, key string) bool {
	return c.storage.Del(ctx, key).Val() == 1
}
