package cache

import (
	"errors"
	"github.com/eko/gocache/lib/v4/cache"
	"go.uber.org/zap"
	"service-api/src/app/helpers"
	"service-api/src/core/system"
)

func Set(key string, val interface{}, exp int) bool {
	return Instance.GetDefaultCacheDrive().Set(key, val, exp)
}

func Get(key string) interface{} {
	return Instance.GetDefaultCacheDrive().Get(key)
}

const (
	REDIS   = "redis"
	MONGODB = "mongodb"
	FILE    = "file"
)

type CacheDrive interface {
	Connect(config interface{}) error
	Set(key string, val interface{}, exp int) bool
	Get(key string) interface{}
	Exists(key string) bool
}

var Instance *CacheManage

type CacheManage struct {
	drive  string
	config system.Cache
	drives map[string]CacheDrive
}

func (c *CacheManage) init() *CacheManage {
	_, err := c.setDefaultDrive(c.drive)
	if err != nil {
		zap.S().Fatalf("register default cache drive fail: %v", err)
	}

	return c
}

func (c *CacheManage) setDefaultDrive(key string) (CacheDrive, error) {
	c.drive = key

	cacheDrive, err := c.Register(key)
	if err != nil {
		return nil, err
	}

	return cacheDrive, nil
}

func (c *CacheManage) GetDefaultCacheDrive() CacheDrive {
	var err error
	if c.drive == "" {
		zap.S().Fatalf("drive not found")
	}

	if c.drives[c.drive] == nil {
		return c.drives[c.drive]
	}

	c.drives[c.drive], err = c.Register(c.drive)
	if err != nil {
		zap.S().Fatalf("register cache drive fail: %v", err)
	}

	return c.drives[c.drive]
}

func (c *CacheManage) Register(key string) (CacheDrive, error) {
	var err error
	if c.drives[key] != nil {
		return c.drives[key], nil
	}
	if key == REDIS {
		c.drives[key], err = c.Connect(key, c.config.Redis)
	} else {
		err = errors.New("register cache drive not exist")
	}

	if err != nil {
		return nil, err
	}

	return c.drives[key], nil
}

func (c *CacheManage) Connect(key string, config interface{}) (CacheDrive, error) {
	var storage CacheDrive
	var err error

	if c.drives[key] != nil {
		return c.drives[key], nil
	}

	switch key {
	case REDIS:
		storage = &CacheRedisDrive{}
		err = storage.Connect(config)

		storage = cache.New[string](storage)
		break
	default:
		storage = nil
		err = errors.New("cache drive not exists")
	}

	return storage, err
}

func (c *CacheManage) GetCacheDriver(key string) (CacheDrive, error) {
	if c.drives[key] != nil {
		return c.drives[key], nil
	}

	return nil, errors.New("cache driver not exists")
}

func NewCacheManageInstance(drive string) CacheManage {
	if helpers.IsEmpty(Instance) {
		Instance = &CacheManage{
			drive:  drive,
			config: system.ConfigInstance.GetCache(),
		}

		Instance.init()
	}

	return *Instance
}
