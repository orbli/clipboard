package storage

import (
	"fmt"

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

func NewStorageRedis(host, port, password string) (*StorageRedisImpl, error) {
	conn := redis.NewClient(
		&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", host, port),
			Password: password,
		},
	)
	_, err := conn.Ping().Result()
	if err != nil {
		return nil, err
	}
	return &StorageRedisImpl{conn}, nil
}

func (s StorageRedisImpl) Get(key string) (Message, error) {
	m := new(Message)
	err := s.c.Get(key).Scan(m)
	if err != nil {
		return Message{}, err
	}
	return *m, nil
}

func (s StorageRedisImpl) Set(key string, value Message) error {
	return s.c.Set(key, value, 0).Err()
}

func (s StorageRedisImpl) Delete(key string) error {
	return s.c.Del(key).Err()
}
