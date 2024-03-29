package cache

import (
	"context"
	"errors"
	"framework/config"
	"framework/log"
	"time"
)

const (
	REDIS   = "redis"
	MONGODB = "mongodb"
	FILE    = "file"
)

type Drive interface {
	Connect(ctx context.Context, cfg map[string]any) error
	SetNx(ctx context.Context, key string, val interface{}, exp time.Duration) (bool, error)
	SetEx(ctx context.Context, key string, val interface{}, exp time.Duration) (string, error)
	Get(ctx context.Context, key string) (string, error)
	Exists(ctx context.Context, key string) bool
	Refresh(ctx context.Context, key string, exp time.Duration) bool
	Del(ctx context.Context, key string) bool
}

type Manage struct {
	drive  string
	drives map[string]Drive
}

func (c *Manage) init() *Manage {

	if c.drives == nil {
		c.drives = map[string]Drive{}
	}

	_, err := c.setDefaultDrive(c.drive)
	if err != nil {
		log.Client.Sugar().Fatalf("register default cache drive fail: %v", err)
	}

	return c
}

func (c *Manage) setDefaultDrive(key string) (Drive, error) {
	c.drive = key

	cacheDrive, err := c.Register(key)
	if err != nil {
		return nil, err
	}

	return cacheDrive, nil
}

func (c *Manage) GetDefaultCacheDrive() Drive {
	var err error
	if c.drive == "" {
		log.Client.Fatal("drive not found")
	}

	if c.drives[c.drive] == nil {
		return c.drives[c.drive]
	}

	c.drives[c.drive], err = c.Register(c.drive)
	if err != nil {
		log.Client.Sugar().Fatalf("register cache drive fail: %v", err)
	}

	return c.drives[c.drive]
}

func (c *Manage) Register(key string) (Drive, error) {
	var err error
	if c.drives[key] != nil {
		return c.drives[key], nil
	}

	conns := config.Client.Get("conns").([]interface{})
	var conn map[string]any
	for _, t := range conns {
		conn = t.(map[string]any)
		if name, ok := conn["driver"].(string); ok && name == key {
			break
		}
	}

	if len(conn) == 0 {
		log.Client.Panic("cache database config conns empty")
	}

	c.drives[key], err = c.Connect(key, conn["driver"].(string), conn)

	if err != nil {
		return nil, err
	}

	return c.drives[key], nil
}

func (c *Manage) Connect(key, tp string, cfg map[string]any) (Drive, error) {
	var storage Drive
	var err error

	if c.drives[key] != nil {
		return c.drives[key], nil
	}

	switch tp {
	case REDIS:
		storage = &RedisDrive{}
		err = storage.Connect(context.Background(), cfg)
		break
	default:
		storage = nil
		err = errors.New("cache drive not exists")
	}

	return storage, err
}

func (c *Manage) GetCacheDriver(key string) (Drive, error) {
	if c.drives[key] != nil {
		return c.drives[key], nil
	}

	return nil, errors.New("cache driver not exists")
}
