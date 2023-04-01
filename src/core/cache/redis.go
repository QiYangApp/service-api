package cache

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"service-api/src/core/system"
)

type CacheRedisMethod interface {
}

type CacheRedisDrive struct {
	storage *redis.Client
	CacheDrive
}

func (c *CacheRedisDrive) Connect(cfg interface{}, cCfg interface{}) error {
	config := cfg.(system.DatabaseRedis)
	cacheConfig := cCfg.(system.CacheRedis)

	c.storage = redis.NewClient(&redis.Options{
		Addr:	  fmt.Sprintf("%s,%d", config.Host, config.Port),
		Password: config.Password, // no password set
		DB:		  cacheConfig.Database,  // use default DB
	})

	c.storage.Ping() {

	}
}
