package storage

import (
	"github.com/go-redis/redis"
)

type (
	StorageRedisImpl struct {
		c *redis.Client
	}
)

var (
	_ StorageStub = StorageRedisImpl{}
)

func NewStorageRedis(addr, password string) (*StorageRedisImpl, error) {
	conn := redis.NewClient(
		&redis.Options{
			Addr:     addr,
			Password: password,
		},
	)
	_, err := conn.Ping().Result()
	if err != nil {
		return nil, err
	}
	return &StorageRedisImpl{conn}, nil
}

func (s StorageRedisImpl) Get(key string) (Value, error) {
	m := new(Value)
	err := s.c.Get(key).Scan(m)
	if err != nil {
		return nil, err
	}
	return *m, nil
}

func (s StorageRedisImpl) Set(value Value) error {
	return s.c.Set(value.Key(), value, 0).Err()
}

func (s StorageRedisImpl) Delete(key string) error {
	return s.c.Del(key).Err()
}

func (s StorageRedisImpl) ListByKey(key string, size int) ([]Value, string, error) {
	panic("I dont think you would use redis for use case accessing this function")
}

func (s StorageRedisImpl) ListByOffset(offset int, size int) ([]Value, error) {
	panic("I dont think you would use redis for use case accessing this function")
}
