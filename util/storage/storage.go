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
	}
)

var (
	Storage StorageStub = StorageGocacheImpl{
		c: cache.New(24*60*time.Minute, 48*60*time.Minute),
	}
	Get       = func(s string) (Value, error) { return Storage.Get(s) }
	Set       = func(v Value) error { return Storage.Set(v) }
	Delete    = func(s string) error { return Storage.Delete(s) }
	ListByKey = func(k string, s int) ([]Value, string, error) { return Storage.ListByKey(k, s) }
)
