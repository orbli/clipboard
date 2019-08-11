package main

import (
	"fmt"
	"os"
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
		storage.Storage = storageRedis
	}
}