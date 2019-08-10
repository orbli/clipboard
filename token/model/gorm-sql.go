package model

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gitlab.com/orbli/clipboard/util/storage"
)

type (
	StorageGormSqlImpl struct {
		db *gorm.DB
	}
	GormToken struct {
		Token
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
	conn.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(GormToken{})
	conn.LogMode(true)
	conn.SetLogger(log.New(os.Stdout, "\n", 0))
	return &StorageGormSqlImpl{conn}, nil
}

func (s StorageGormSqlImpl) Get(key string) (storage.Value, error) {
	rt := new(GormToken)
	where := &GormToken{Token: Token{Token: []byte(key)}}
	if err := s.db.Where(where).First(rt).Error; err != nil {
		return Token{}, err
	}
	return rt.Token, nil
}

func (s StorageGormSqlImpl) Set(value storage.Value) error {
	v := &GormToken{Token: value.(Token)}
	return s.db.Model(&GormToken{}).Updates(v).Error
}

func (s StorageGormSqlImpl) Delete(key string) error {
	where := &GormToken{Token: Token{Token: []byte(key)}}
	return s.db.Where(where).Delete(&GormToken{}).Error
}

func (s StorageGormSqlImpl) ListByKey(key string, size int) ([]storage.Value, string, error) {
	panic("Not yet implement je")
}

func (s StorageGormSqlImpl) ListByOffset(offset int, size int) ([]storage.Value, error) {
	panic("Not yet implement je")
}
