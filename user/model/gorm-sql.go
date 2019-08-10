package model

import (
	"log"
	"os"

	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gitlab.com/orbli/clipboard/util/storage"
)

type (
	StorageGormSqlImpl struct {
		db *gorm.DB
	}
	GormUser struct {
		User
	}
)

var (
	_ storage.StorageStub = StorageGormSqlImpl{}
)

func NewStorageGormSql(dsn string) (*StorageGormSqlImpl, error) {
	conn, err := gorm.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	conn.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(GormUser{})
	conn.LogMode(true)
	conn.SetLogger(log.New(os.Stdout, "\n", 0))
	return &StorageGormSqlImpl{conn}, nil
}

func (s StorageGormSqlImpl) Get(key string) (storage.Value, error) {
	keyuint64, err := strconv.ParseUint(key, 10, 64)
	if err != nil {
		return nil, err
	}
	rt := new(GormUser)
	where := &GormUser{User: User{Id: keyuint64}}
	if err := s.db.Where(where).First(rt).Error; err != nil {
		return User{}, err
	}
	return rt.User, nil
}

func (s StorageGormSqlImpl) Set(value storage.Value) error {
	v := &GormUser{User: value.(User)}
	return s.db.Model(&GormUser{}).Updates(v).Error
}

func (s StorageGormSqlImpl) Delete(key string) error {
	keyuint64, err := strconv.ParseUint(key, 10, 64)
	if err != nil {
		return err
	}
	where := &GormUser{User: User{Id: keyuint64}}
	return s.db.Where(where).Delete(&GormUser{}).Error
}

func (s StorageGormSqlImpl) ListByKey(key string, size int) ([]storage.Value, string, error) {
	panic("Not yet implement je")
}

func (s StorageGormSqlImpl) ListByOffset(offset int, size int) ([]storage.Value, error) {
	panic("Not yet implement je")
}
