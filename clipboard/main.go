package main

import (
	"context"
	"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/memcache"
	"io/ioutil"
	"log"
	"net/http"
)

type (
	Handler struct{}
)

func main() {
	log.Printf("Listening through appengine")
	http.HandleFunc("/", Handler{}.ServeHTTP)
	appengine.Main()
}
func (Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	bodybytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	body := string(bodybytes)
	log.Printf("Method: %s , URI: %s, Body: %s", r.Method, r.URL.Path, body)
	switch r.Method {
	case "GET":
		fmt.Fprint(w, GetValue(ctx, r.URL.Path))
	case "POST":
		fmt.Fprint(w, SetValue(ctx, r.URL.Path, body))
	case "DELETE":
		fmt.Fprint(w, DelValue(ctx, r.URL.Path))
	default:
		http.NotFound(w, r)
	}
	return
}
func GetValue(ctx context.Context, key string) string {
	if item, err := memcache.Get(ctx, key); err == memcache.ErrCacheMiss {
		return "Error cache miss"
	} else if err != nil {
		return err.Error()
	} else {
		return string(item.Value)
	}
}
func SetValue(ctx context.Context, key, value string) string {
	item := &memcache.Item{
		Key:   key,
		Value: []byte(value),
	}
	if err := memcache.Set(ctx, item); err != nil {
		return err.Error()
	} else {
		return GetValue(ctx, key)
	}
}
func DelValue(ctx context.Context, key string) string {
	if err := memcache.Delete(ctx, key); err != nil {
		return err.Error()
	} else {
		return ""
	}
}
