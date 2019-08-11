package main

import (
	"fmt"
	"gitlab.com/orbli/clipboard/user/model"
	"gitlab.com/orbli/clipboard/util/storage"
	"os"
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
		storage.Storage = storageGormSql
	}
}
