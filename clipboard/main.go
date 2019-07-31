package main

import (
	"clipboard/util"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type (
	Handler struct{}
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	httpServer := http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: Handler{},
	}
	log.Printf("Serving http")
	log.Fatal(httpServer.ListenAndServe())
}

func (Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	bodybytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}
	key := strings.Trim(r.URL.Path, "/")
	value := string(bodybytes)
	log.Printf("Method: %s , Key: %s, Balue: %s", r.Method, key, value)
	switch r.Method {
	case "GET":
		fmt.Fprint(w, Read(key))
	case "POST":
		Create(key, value)
	case "PUT":
		Update(key, value)
	case "DELETE":
		Delete(key)
	default:
		http.NotFound(w, r)
	}
	return
}

func Create(key, value string) {
	SetValue(key, value)
}
func Read(key string) string {
	return GetValue(key)
}
func Update(key, value string) {
	SetValue(key, value)
}
func Delete(key string) {
	DelValue(key)
}

var (
	colref = util.FirestoreClient.Collection("clipboards")
)

func GetValue(key string) string {
	data, err := colref.Doc(key).Get(util.FirestoreContext)
	if err != nil {
		panic(err)
	}
	return data.Data()["msg"].(string)
}
func SetValue(key, value string) {
	rt := map[string]interface{}{
		"msg": value,
	}
	_, err := colref.Doc(key).Set(util.FirestoreContext, rt)
	if err != nil {
		panic(err)
	}
}
func DelValue(key string) {
	_, err := colref.Doc(key).Delete(util.FirestoreContext)
	if err != nil {
		panic(err)
	}
}
