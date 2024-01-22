package cache

import (
	"context"
	"errors"
	"service-api/src/core/config"
	"service-api/src/core/logger"
	"time"
)

const (
	REDIS   = "redis"
	MONGODB = "mongodb"
	FILE    = "file"
)

type CacheDrive interface {
	Connect(ctx context.Context, cfg interface{}, cCfg interface{}) error
	SetNx(ctx context.Context, key string, val interface{}, exp time.Duration) (bool, error)
	SetEx(ctx context.Context, key string, val interface{}, exp time.Duration) (string, error)
	Get(ctx context.Context, key string) (string, error)
	Exists(ctx context.Context, key string) bool
	Refresh(ctx context.Context, key string, exp time.Duration) bool
	Del(ctx context.Context, key string) bool
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
		err = storage.Connect(context.TODO(), cfg, cCfg)
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
