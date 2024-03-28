package captcha

import (
	"context"
	"framework/cache"
	"framework/exceptions"
	"service-api/resources/lang"
	"sync"
	"time"
)

type StoreValue struct {
	sync.RWMutex
	store cache.Drive
	exp   time.Duration
}

func (i *StoreValue) Set(id string, value string) error {
	id = i.getCacheKey(id)

	if _, err := i.store.SetEx(context.Background(), id, value, i.exp); err != nil {
		return exceptions.New(lang.CaptchaErrorStoreCode)
	}

	return nil
}

func (i *StoreValue) Get(id string, clear bool) (value string) {
	i.Lock()
	defer i.Unlock()

	id = i.getCacheKey(id)

	if !i.store.Exists(context.Background(), id) {
		return
	}

	value, err := i.store.Get(context.Background(), id)
	if err == nil && clear {
		i.store.Del(context.Background(), id)
	}

	return
}

func (i *StoreValue) Verify(id string, answer any, clear bool) bool {
	res := i.Get(i.getCacheKey(id), clear)

	return res != "" && key == answer
}

func (i *StoreValue) getCacheKey(id string) string {
	return "captcha-" + id
}

func NewStoreValue(store cache.Drive, t time.Duration) Store {
	return &StoreValue{store: store, exp: t}
}
