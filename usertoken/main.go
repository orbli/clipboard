package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	"gitlab.com/orbli/clipboard/usertoken/handler"
	pb "gitlab.com/orbli/clipboard/usertoken/proto"
)

func main() {
	service := micro.NewService(
		micro.Name("orbli.micro.usertoken"),
		micro.Version("latest"),
	)
	service.Init()
	pb.RegisterUsertokenServiceHandler(service.Server(), &handler.UsertokenService{service.Client()})
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
