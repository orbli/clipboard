package main

import (

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	"gitlab.com/orbli/clipboard/clipboard/handler"
	pb "gitlab.com/orbli/clipboard/clipboard/proto"
)

func main() {
	service := micro.NewService(
		micro.Name("orbli.micro.clipboard"),
		micro.Version("latest"),
	)
	service.Init()
	pb.RegisterClipboardHandler(service.Server(), new(handler.Clipboard))
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
