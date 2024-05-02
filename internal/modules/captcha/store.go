package captcha

import (
	"errors"
	"frame/modules/cache"
	"frame/util/optional"
	"service-api/resources/i18n"
	"time"
)

type StoreValue[T any] struct {
	store *cache.Operation[T]
	exp   time.Duration
}

func (i *StoreValue[T]) Set(id string, value any) error {
	id = i.getCacheKey(id)

	if i.store.SetEx(id, value, i.exp) == false {
		return errors.New(i18n.CaptchaErrorStoreCode)
	}

	return nil
}

func (i *StoreValue[T]) Get(key string, clear bool) optional.Option[T] {
	var value optional.Option[T]

	key = i.getCacheKey(key)

	if !i.store.Exists(key) {
		return optional.None[T]()
	}

	value, err := i.store.Get(key)
	if err == nil && clear {
		i.store.Del(key)
	}

	return value
}

func (i *StoreValue[T]) Exist(key string) bool {
	key = i.getCacheKey(key)
	if !i.store.Exists(key) {
		return false
	}

	return true
}

func (i *StoreValue[T]) Del(key string) bool {
	key = i.getCacheKey(key)
	if !i.store.Exists(key) {
		return false
	}

	return i.store.Del(key)
}

func (i *StoreValue[T]) getCacheKey(id string) string {
	return "captcha-" + id
}

func NewStoreValue[T any](store *cache.Operation[T], t time.Duration) Store[T] {
	return &StoreValue[T]{store: store, exp: t}
}
