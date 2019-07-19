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
        log.Printf("Body: %s", body)
        switch r.Method {
        case "GET":
                if r.URL.Path == "/" {
                        b, _ := ioutil.ReadFile("index.html")
                        fmt.Fprint(w, string(b))
                        return
                } else {
                        fmt.Fprint(w, GetValue(ctx, r.URL.Path))
                }
        case "POST":
                if r.URL.Path != "/" {
                        fmt.Fprint(w, SetValue(ctx, r.URL.Path, body))
                }
        case "DELETE":
                if r.URL.Path != "/" {
                        if err := memcache.Delete(ctx, r.URL.Path); err != nil {
                                fmt.Fprint(w, err.Error())
                        }
                }
        default:
                http.NotFound(w, r)
                return
        }
        w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
        w.Header().Set("Pragma", "no-cache")
        w.Header().Set("Expires", "0")
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
