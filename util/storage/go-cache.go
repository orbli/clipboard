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

func (s StorageGocacheImpl) Get(key string) (Value, error) {
	if value, found := s.c.Get(key); found {
		return value.(Value), nil
	}
	return nil, errors.New("Not found")
}

func (s StorageGocacheImpl) Set(value Value) error {
	s.c.Set(value.Key(), value, cache.DefaultExpiration)
	return nil
}

func (s StorageGocacheImpl) Delete(key string) error {
	s.c.Delete(key)
	return nil
}
