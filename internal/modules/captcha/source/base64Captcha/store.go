package base64Captcha

import (
	"framework/cache"
	"framework/exceptions"
	"service-api/resources/lang"
	"sync"
	"time"
)

type Store struct {
	sync.RWMutex
	exp time.Duration
}

func (i *Store) Set(id string, value string) error {
	if !cache.SetEx(id, value, i.exp) {
		return exceptions.New(lang.CaptchaErrorStoreCode)
	}

	return nil
}

func (i *Store) Get(id string, clear bool) (value string) {
	i.Lock()
	defer i.Unlock()

	if !cache.Exists(id) {
		return
	}

	err := cache.Get(id, value)
	if err == nil && clear {
		cache.Del(id)
	}

	return
}

func (i *Store) Verify(id string, answer string, clear bool) bool {
	key := i.Get(id, clear)

	return key != "" && key == answer
}

func NewStore(t time.Duration) *Store {
	return &Store{exp: t}
}
