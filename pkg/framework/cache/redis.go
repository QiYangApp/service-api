package cache

import (
	"context"
	"framework/log"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisMethod interface {
}

type RedisDrive struct {
	storage *redis.Client
	Drive
}

func (c *RedisDrive) Connect(ctx context.Context, cfg map[string]any) error {
	ctx = context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}
	//
	//err := c.storage.Set(ctx, "key", "value", 0).Err()
	//if err != nil {
	//	panic(err)
	//}
	//
	//_, err = c.storage.Ping(ctx).Result()
	//if err != nil {
	//	return err
	//}

	return nil
}

func (c *RedisDrive) Get(ctx context.Context, key string) (string, error) {
	return c.storage.Get(ctx, key).Result()
}

func (c *RedisDrive) SetNx(ctx context.Context, key string, val interface{}, exp time.Duration) (bool, error) {
	return c.storage.SetNX(ctx, key, val, exp).Result()
}

func (c *RedisDrive) SetEx(ctx context.Context, key string, val interface{}, exp time.Duration) (string, error) {
	return c.storage.SetEx(ctx, key, val, exp).Result()
}

func (c *RedisDrive) Exists(ctx context.Context, key string) bool {
	state := c.storage.Exists(ctx, key)
	if state.Err() != nil {
		return false
	}

	return state.Val() == 1
}

func (c *RedisDrive) Refresh(ctx context.Context, key string, exp time.Duration) bool {

	if c.Exists(ctx, key) == false {
		return false
	}

	valState, err := c.Get(ctx, key)
	if err != nil {
		log.Client().Sugar().Warnf("cache drive refresh fail, get origin data fail, key, %s", err)
		return false
	}

	state, err := c.SetNx(ctx, key, valState, exp)
	if err != nil {
		log.Client().Sugar().Warnf("cache drive refresh fail, set new data fail, key, %s", err)
		return false
	}

	return state
}

func (c *RedisDrive) Del(ctx context.Context, key string) bool {
	return c.storage.Del(ctx, key).Val() == 1
}
