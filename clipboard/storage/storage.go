package storage

import (
	"time"

	"github.com/patrickmn/go-cache"
)

type (
	StorageStub interface {
		Get(key string) (Message, error)
		Set(key string, value Message) error
		Delete(key string) error
	}
)

var (
	Storage StorageStub = StorageGocacheImpl{
		c: cache.New(24*60*time.Minute, 48*60*time.Minute),
	}
	Get    func(string) (Message, error) = Storage.Get
	Set    func(string, Message) error   = Storage.Set
	Delete func(string) error            = Storage.Delete
)

func Replace(newStorage StorageStub) {
	Get = newStorage.Get
	Set = newStorage.Set
	Delete = newStorage.Delete
}
