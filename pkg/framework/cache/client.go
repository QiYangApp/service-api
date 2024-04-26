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

	if len(c.drive) == 0 {
		return c
	}

	if err := c.SetDefaultDrive(c.drive); err != nil {
		log.Client.Sugar().Fatalf("register default cache drive fail: %v", err)
	}

	return c
}

func (c *Manage) SetDefaultDrive(key string) error {
	c.drive = key

	_, err := c.Register(key)
	if err != nil {
		log.Client.Sugar().Warnf("set default cache dirver fail, err: %v", err)
		return err
	}

	return nil
}

func (c *Manage) GetDefaultDrive() Drive {
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

	conn := config.Client.GetStringMap("conns." + key)
	if len(conn) == 0 {
		log.Client.Panic("cache database config conns empty")
		return nil, errors.New("cache database config conns empty")
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
