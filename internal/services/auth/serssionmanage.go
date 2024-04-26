package auth

import (
	"framework/cache"
	"framework/utils/optional"
)

type TokenStore struct {
	Token string
}

func (t *TokenStore) client() *cache.Operation[map[string]any] {
	return cache.NewOperation[map[string]any](cache.Client.GetDefaultDrive())
}

func (t *TokenStore) context() optional.Option[map[string]any] {
	var data
	if data, err := t.client().Get(t.Token); err != nil {
		return optional.None[map[string]any]()
	}

	return
}

func (t *TokenStore) Set(key, val any) error {

}

func (t *TokenStore) Get(key any) any {

}