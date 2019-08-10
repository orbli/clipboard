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

func (s StorageGocacheImpl) ListByKey(key string, size int) ([]Value, string, error) {
	rt := []Value{}
	for k, v := range s.c.Items() {
		if k < key {
			continue
		}
		if size == 0 {
			return rt, k, nil
		}
		rt = append(rt, v.Object.(Value))
		size -= 1
	}
	return rt, "", nil
}
func (s StorageGocacheImpl) ListByOffset(offset int, size int) ([]Value, error) {
	rt := []Value{}
	for _, v := range s.c.Items() {
		if offset != 0 {
			offset -= 1
			continue
		}
		if size == 0 {
			return rt, nil
		}
		rt = append(rt, v.Object.(Value))
		size -= 1
	}
	return rt, nil
}
