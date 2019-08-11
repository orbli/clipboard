package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-micro/util/log"
	"gitlab.com/orbli/clipboard/usertoken/handler"
	pb "gitlab.com/orbli/clipboard/usertoken/proto"
	"gitlab.com/orbli/clipboard/usertoken/subscriber"
)

func main() {
	service := micro.NewService(
		micro.Name("orbli.micro.usertoken"),
		micro.Version("latest"),
	)
	service.Init()

	micro.RegisterSubscriber(
		"orbli.micro.user",
		service.Server(),
		&subscriber.UserSubscriber{service.Client()},
		server.SubscriberQueue("orbli.micro.usertoken"),
	)
	pb.RegisterUsertokenServiceHandler(
		service.Server(),
		&handler.UsertokenService{service.Client()},
	)
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
