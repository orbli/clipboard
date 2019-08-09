package main

import (
	"fmt"
	"os"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	"gitlab.com/orbli/clipboard/token/handler"
	"gitlab.com/orbli/clipboard/token/model"
	pb "gitlab.com/orbli/clipboard/token/proto"
	"gitlab.com/orbli/clipboard/util/storage"
)

func init() {
	if os.Getenv("USE_REDIS") == "1" {
		storageRedis, err := storage.NewStorageRedis(
			fmt.Sprintf("%s:%s",
				os.Getenv("REDIS_HOST"),
				os.Getenv("REDIS_PORT"),
			),
			os.Getenv("REDIS_PSWD"),
		)
		if err != nil {
			panic(err)
		}
		storage.Replace(storageRedis)
	}
	if os.Getenv("USE_SQL") == "1" {
		storageGormSql, err := model.NewStorageGormSql(
			fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
				os.Getenv("SQL_USER"),
				os.Getenv("SQL_PSWD"),
				os.Getenv("SQL_HOST"),
				os.Getenv("SQL_PORT"),
				os.Getenv("SQL_DBSE"),
			),
		)
		if err != nil {
			panic(err)
		}
		storage.Replace(storageGormSql)
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
