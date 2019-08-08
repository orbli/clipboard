package storage

import (
	"time"

	"github.com/patrickmn/go-cache"
)

type (
	StorageStub interface {
		Get(key string) (Token, error)
		Set(key string, value Token) error
		Delete(key string) error
	}
)

var (
	Storage StorageStub = StorageGocacheImpl{
		c: cache.New(24*60*time.Minute, 48*60*time.Minute),
	}
	Get    func(string) (Token, error) = Storage.Get
	Set    func(string, Token) error   = Storage.Set
	Delete func(string) error          = Storage.Delete
)

func Replace(newStorage StorageStub) {
	Get = newStorage.Get
	Set = newStorage.Set
	Delete = newStorage.Delete
}
