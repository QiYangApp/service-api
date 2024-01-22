package cache

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"service-api/src/core/config"
	"service-api/src/core/helpers"
	"service-api/src/core/logger"
	"sync"
	"time"
)

var Instance *CacheManage
var once = sync.Once{}

func SetEx(key string, val interface{}, exp time.Duration) bool {
	payload, SerializeErr := helpers.Serialize(val)
	if SerializeErr != nil {
		logger.S().Warnw(
			"cache",
			zap.String("mes", "cache drive SerializeErr key error"),
			zap.String("key", key),
			zap.Any("val", val),
		)
		return false
	}

	state, err := Instance.GetDefaultCacheDrive().SetEx(context.TODO(), key, payload, exp)
	if err != nil {
		logger.S().Infof("cache drive set key error, key: %s, val: %v", key, payload)
		return false
	}

	return state != ""
}

func SetNx(key string, val interface{}, exp time.Duration) bool {
	payload, SerializeErr := helpers.Serialize(val)
	if SerializeErr != nil {
		logger.S().Warnw(
			"cache",
			zap.String("mes", "cache drive SerializeErr key error"),
			zap.String("key", key),
			zap.Any("val", val),
		)
		return false
	}

	state, err := Instance.GetDefaultCacheDrive().SetNx(context.TODO(), key, payload, exp)
	if err != nil {
		logger.S().Infof("cache drive set key error, key: %s, val: %v", key, payload)
		return false
	}

	return state
}

func Get(key string, ptrValue interface{}) error {
	payload, err := Instance.GetDefaultCacheDrive().Get(context.Background(), key)
	if errors.Is(err, redis.Nil) {
		return errors.New("data is empty")
	}

	if err != nil {
		return err
	}

	return helpers.Deserialize([]byte(payload), ptrValue)
}

func Exists(key string) bool {
	return Instance.GetDefaultCacheDrive().Exists(context.TODO(), key)
}

func Refresh(key string, exp time.Duration) bool {
	return Instance.GetDefaultCacheDrive().Refresh(context.TODO(), key, exp)
}

func Del(key string) bool {
	return Instance.GetDefaultCacheDrive().Del(context.TODO(), key)
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
