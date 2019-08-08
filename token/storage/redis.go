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

func (s StorageRedisImpl) Get(key string) (Token, error) {
	m := new(Token)
	err := s.c.Get(key).Scan(m)
	if err != nil {
		return Token{}, err
	}
	return *m, nil
}

func (s StorageRedisImpl) Set(key string, value Token) error {
	return s.c.Set(key, value, 0).Err()
}

func (s StorageRedisImpl) Delete(key string) error {
	return s.c.Del(key).Err()
}
