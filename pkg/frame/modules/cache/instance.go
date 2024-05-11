package cache

import (
	"frame/conf"
	"frame/util/optional"
	"sync"
	"time"
)

var singleton *Manage
var once = sync.Once{}

func Client() *Manage {
	once.Do(func() {
		singleton = &Manage{
			drive: conf.Client().GetString("cache.driver"),
		}

		singleton.init()
	})

	return singleton
}

func SetEx(key string, val any, exp time.Duration) bool {
	return NewOperation[any](Client().GetDefaultDrive()).SetEx(key, val, exp)
}

func SetNx(key string, val any, exp time.Duration) bool {
	return NewOperation[any](Client().GetDefaultDrive()).SetNx(key, val, exp)
}

func Get[T any](key string) (optional.Option[T], error) {
	return NewOperation[T](Client().GetDefaultDrive()).Get(key)
}

func Exists(key string) bool {
	return NewOperation[any](Client().GetDefaultDrive()).Exists(key)
}

func Refresh(key string, exp time.Duration) bool {
	return NewOperation[any](Client().GetDefaultDrive()).Refresh(key, exp)
}

func Del(key string) bool {
	return NewOperation[any](Client().GetDefaultDrive()).Del(key)
}
