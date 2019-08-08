package storage

import (
	"errors"

	"github.com/patrickmn/go-cache"
)

type (
	StorageGocacheImpl struct {
		c *cache.Cache
	}
)

func (s StorageGocacheImpl) Get(key string) (Token, error) {
	if value, found := s.c.Get(key); found {
		return value.(Token), nil
	}
	return Token{}, errors.New("Not found")
}

func (s StorageGocacheImpl) Set(key string, value Token) error {
	s.c.Set(key, value, cache.DefaultExpiration)
	return nil
}

func (s StorageGocacheImpl) Delete(key string) error {
	s.c.Delete(key)
	return nil
}
