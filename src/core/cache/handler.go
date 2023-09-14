package cache

import (
	"errors"
	"service-api/src/core/config"
	"service-api/src/core/logger"
)

const (
	REDIS   = "redis"
	MONGODB = "mongodb"
	FILE    = "file"
)

type CacheDrive interface {
	Connect(cfg interface{}, cCfg interface{}) error
	SetNx(key string, val interface{}, exp int) (bool, error)
	SetEx(key string, val interface{}, exp int) (string, error)
	Get(key string) (string, error)
	Exists(key string) bool
	Refresh(key string, exp int) bool
	Del(key string) bool
}

type CacheManage struct {
	drive          string
	cacheConfig    config.Cache
	databaseConfig config.Database
	drives         map[string]CacheDrive
}

func (c *CacheManage) init() *CacheManage {

	if c.drives == nil {
		c.drives = map[string]CacheDrive{}
	}

	_, err := c.setDefaultDrive(c.drive)
	if err != nil {
		logger.S().Fatalf("register default cache drive fail: %v", err)
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
		logger.S().Fatalf("drive not found")
	}

	if c.drives[c.drive] == nil {
		return c.drives[c.drive]
	}

	c.drives[c.drive], err = c.Register(c.drive)
	if err != nil {
		logger.S().Fatalf("register cache drive fail: %v", err)
	}

	return c.drives[c.drive]
}

func (c *CacheManage) Register(key string) (CacheDrive, error) {
	var err error
	if c.drives[key] != nil {
		return c.drives[key], nil
	}
	if key == REDIS {
		c.drives[key], err = c.Connect(key, c.databaseConfig.Cache, c.cacheConfig.Redis)
	} else {
		err = errors.New("register cache drive not exist")
	}

	if err != nil {
		return nil, err
	}

	return c.drives[key], nil
}

func (c *CacheManage) Connect(key string, cfg interface{}, cCfg interface{}) (CacheDrive, error) {
	var storage CacheDrive
	var err error

	if c.drives[key] != nil {
		return c.drives[key], nil
	}

	switch key {
	case REDIS:
		storage = &CacheRedisDrive{}
		err = storage.Connect(cfg, cCfg)
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
