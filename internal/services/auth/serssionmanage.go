package auth

import (
	"framework/cache"
)

type SessionClient struct {
	Token string
}

func (s *SessionClient) client() cache.Drive {
	return cache.Client.GetDefaultDrive()
}

func (s *SessionClient) Set(key string, val any) error {
}

func (s *SessionClient) Get(key string, val any) error {

}

func (s *SessionClient) Del(key ...string) error {

}

func (s *SessionClient) Update(data map[string]any) error {

}
