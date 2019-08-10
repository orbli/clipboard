package storage

import (
	"time"

	"github.com/patrickmn/go-cache"
)

type (
	StorageStub interface {
		Get(key string) (Value, error)
		Set(Value) error
		Delete(key string) error

		ListByKey(key string, size int) ([]Value, string, error)
		ListByOffset(offset int, size int) ([]Value, error)
	}
)

var (
	Storage StorageStub = StorageGocacheImpl{
		c: cache.New(24*60*time.Minute, 48*60*time.Minute),
	}
	Get          func(string) (Value, error)                = Storage.Get
	Set          func(Value) error                          = Storage.Set
	Delete       func(string) error                         = Storage.Delete
	ListByKey    func(string, int) ([]Value, string, error) = Storage.ListByKey
	ListByOffset func(int, int) ([]Value, error)            = Storage.ListByOffset
)

func Replace(newStorage StorageStub) {
	Get = newStorage.Get
	Set = newStorage.Set
	Delete = newStorage.Delete
	ListByKey = newStorage.ListByKey
	ListByOffset = newStorage.ListByOffset
}
