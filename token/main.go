package main

import (
	"os"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	"gitlab.com/orbli/clipboard/token/handler"
	pb "gitlab.com/orbli/clipboard/token/proto"
	"gitlab.com/orbli/clipboard/token/storage"
)

func init() {
	if os.Getenv("USE_REDIS") == "1" {
		host := os.Getenv("REDIS_HOST")
		port := os.Getenv("REDIS_PORT")
		pswd := os.Getenv("REDIS_PSWD")
		log.Logf("Use redis: %s:%s %s", host, port, pswd)
		storageRedis, err := storage.NewStorageRedis(host, port, pswd)
		if err != nil {
			panic(err)
		}
		storage.Replace(storageRedis)
	}
}

func main() {
	service := micro.NewService(
		micro.Name("orbli.micro.token"),
		micro.Version("latest"),
	)
	service.Init()
	pb.RegisterTokenServiceHandler(service.Server(), new(handler.TokenService))
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
