package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	"gitlab.com/orbli/clipboard/token/handler"
	pb "gitlab.com/orbli/clipboard/token/proto"
)

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
