package cache

import (
	"context"
	"service-api/src/core/config"
	"service-api/src/core/logger"
	"sync"
)

var ctx = context.Background()
var Instance *CacheManage
var once = sync.Once{}

func SetEx(key string, val interface{}, exp int) bool {
	state, err := Instance.GetDefaultCacheDrive().SetEx(key, val, exp)
	if err != nil {
		logger.S().Infof("cache drive set key error, key: %s, val: %v", key, val)
		return false
	}

	return state != ""
}

func SetNx(key string, val interface{}, exp int) bool {
	state, err := Instance.GetDefaultCacheDrive().SetNx(key, val, exp)
	if err != nil {
		logger.S().Infof("cache drive set key error, key: %s, val: %v", key, val)
		return false
	}

	return state
}

func Get(key string) (string, error) {
	return Instance.GetDefaultCacheDrive().Get(key)
}

func Exists(key string) bool {
	return Instance.GetDefaultCacheDrive().Exists(key)
}

func Refresh(key string, exp int) bool {
	return Instance.GetDefaultCacheDrive().Refresh(key, exp)
}

func Del(key string) bool {
	return Instance.GetDefaultCacheDrive().Del(key)
}

func NewInstance(drive string) CacheManage {
	once.Do(func() {
		Instance = &CacheManage{
			drive:          drive,
			cacheConfig:    config.Instance.GetCache(),
			databaseConfig: config.Instance.GetDatabase(),
		}

		Instance.init()
	})

	return *Instance
}
