package cache

import (
	"context"
	"errors"
	"framework/log"
	"framework/utils"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"sync"
	"time"
)

var singleton *Manage
var once = sync.Once{}

var Client Drive = nil

func NewInstance(drive string) Manage {
	once.Do(func() {
		singleton = &Manage{
			drive: drive,
		}

		singleton.init()

		Client = singleton.GetDefaultCacheDrive()
	})

	return *singleton
}

func Instance() *Manage {
	return singleton
}

func SetEx(key string, val interface{}, exp time.Duration) bool {
	payload, SerializeErr := utils.Serialize(val)
	if SerializeErr != nil {
		log.Client.Warn(
			"cache",
			zap.String("mes", "cache drive SerializeErr key error"),
			zap.String("key", key),
			zap.Any("val", val),
		)
		return false
	}

	state, err := singleton.GetDefaultCacheDrive().SetEx(context.TODO(), key, payload, exp)
	if err != nil {
		log.Client.Sugar().Info("cache drive set key error, key: %s, val: %v", key, payload)
		return false
	}

	return state != ""
}

func SetNx(key string, val interface{}, exp time.Duration) bool {
	payload, SerializeErr := utils.Serialize(val)
	if SerializeErr != nil {
		log.Client.Warn(
			"cache",
			zap.String("mes", "cache drive SerializeErr key error"),
			zap.String("key", key),
			zap.Any("val", val),
		)
		return false
	}

	state, err := singleton.GetDefaultCacheDrive().SetNx(context.TODO(), key, payload, exp)
	if err != nil {
		log.Client.Sugar().Infof("cache drive set key error, key: %s, val: %v", key, payload)
		return false
	}

	return state
}

func Get(key string, ptrValue interface{}) error {
	payload, err := singleton.GetDefaultCacheDrive().Get(context.Background(), key)
	if errors.Is(err, redis.Nil) {
		return errors.New("data is empty")
	}

	if err != nil {
		return err
	}

	return utils.Deserialize([]byte(payload), ptrValue)
}

func Exists(key string) bool {
	return singleton.GetDefaultCacheDrive().Exists(context.TODO(), key)
}

func Refresh(key string, exp time.Duration) bool {
	return singleton.GetDefaultCacheDrive().Refresh(context.TODO(), key, exp)
}

func Del(key string) bool {
	return singleton.GetDefaultCacheDrive().Del(context.TODO(), key)
}
