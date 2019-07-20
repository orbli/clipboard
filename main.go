package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type (
	Handler struct{}
)

var (
	cache *redis.Client
)

func main() {
	cache = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PSWD"),
		DB:       0,
	})
	if _, err := cache.Ping().Result(); err != nil {
		log.Fatal(err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	httpServer := http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: Handler{},
	}
	log.Println("Serve")
	log.Fatal(httpServer.ListenAndServe())
}
func (Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
			fmt.Fprint(w, GetValue(r.URL.Path))
		}
	case "POST":
		if r.URL.Path != "/" {
			fmt.Fprint(w, SetValue(r.URL.Path, body))
		}
	case "DELETE":
		if r.URL.Path != "/" {
			fmt.Fprint(w, DelValue(r.URL.Path))
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
func GetValue(key string) string {
	if val, err := cache.Get(key).Result(); err != nil {
		return err.Error()
	} else {
		return fmt.Sprintf("%v", val)
	}
}
func SetValue(key, value string) string {
	if err := cache.Set(key, value, 0).Err(); err != nil {
		return err.Error()
	} else {
		return GetValue(key)
	}
}
func DelValue(key string) string {
	if err := cache.Del(key).Err(); err != nil {
		return err.Error()
	} else {
		return GetValue(key)
	}
}
