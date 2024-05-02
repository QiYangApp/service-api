package auth

import "frame/modules/cache"

type SessionClient struct {
	Token       string
	UserId      int64
	RoleId      int64
	IsSigned    bool
	IsTwoFa     bool // 双重
	IsRegister  bool
	Language    string
	Theme       string
	Ext         map[string]any
	StoreClient cache.Operation[map[string]any]
}

func (s *SessionClient) Set(key string, val any) *SessionClient {
	s.Ext[key] = val

	return s
}

func (s *SessionClient) Get(key string, val any) error {
	return nil
}

func (s *SessionClient) Del(key ...string) error {
	return nil
}

func (s *SessionClient) Update(data map[string]any) error {
	return nil
}
