package cache

import (
	"errors"
	"github.com/amanbolat/umsystem/internal/auth"
	"github.com/amanbolat/umsystem/pkg"
)

var ErrTokenAlreadyExists = errors.New("auth token with given key does already exist")
var ErrTokenNotFound = errors.New("token with given key was not found")

type InMemoryTokenCache struct {
	m *pkg.Map[auth.AuthToken]
}

func NewInMemoryTokenCache() *InMemoryTokenCache {
	return &InMemoryTokenCache{
		m: pkg.NewMap[auth.AuthToken](),
	}
}

func (i *InMemoryTokenCache) SaveToken(token auth.AuthToken) error {
	if i.m.Exists(token.Key) {
		return ErrTokenAlreadyExists
	}
	i.m.Set(token.Key, token)
	return nil
}

func (i *InMemoryTokenCache) DeleteToken(key string) error {
	i.m.Delete(key)
	return nil
}

func (i *InMemoryTokenCache) GetToken(key string) (auth.AuthToken, error) {
	val, ok := i.m.Get(key)
	if !ok {
		return auth.AuthToken{}, ErrTokenNotFound
	}

	return val, nil
}
