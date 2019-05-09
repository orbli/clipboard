package main

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type (
	Handler struct{}
)

var (
	c *cache.Cache
)

func main() {
	c = cache.New(24*60*time.Minute, 48*60*time.Minute)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Printf("Listening on port %s", port)
	httpServer := http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: Handler{},
	}
	log.Fatal(httpServer.ListenAndServe())
}
func (Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body := BodyString(r)
	log.Printf("Method: %s , URI: %s, Body: %s", r.Method, r.URL.Path, body)
	switch r.Method {
	case "GET":
		if r.URL.Path != "/" {
			fmt.Fprint(w, GetValue(c, r.URL.Path))
		} else {
			b, _ := ioutil.ReadFile("index.html")
			fmt.Fprint(w, string(b))
		}
	case "POST":
		if r.URL.Path != "/" {
			c.Set(r.URL.Path, body, cache.DefaultExpiration)
			fmt.Fprint(w, GetValue(c, r.URL.Path))
		}
	case "DELETE":
		if r.URL.Path != "/" {
			c.Set(r.URL.Path, "", cache.DefaultExpiration)
			fmt.Fprint(w, GetValue(c, r.URL.Path))
		}
	default:
		http.NotFound(w, r)
	}
	return
}
func GetValue(c *cache.Cache, key string) string {
	if value, found := c.Get(key); found {
		return value.(string)
	} else {
		return ""
	}
}
func BodyString(r *http.Request) string {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	return string(body)
}
