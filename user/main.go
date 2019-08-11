package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	"gitlab.com/orbli/clipboard/user/handler"
	pb "gitlab.com/orbli/clipboard/user/proto"
)

func main() {
	service := micro.NewService(
		micro.Name("orbli.micro.user"),
		micro.Version("latest"),
	)
	service.Init()
	pb.RegisterUserServiceHandler(
		service.Server(),
		&handler.UserService{
			micro.NewPublisher("orbli.micro.user", service.Client()),
		},
	)
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
