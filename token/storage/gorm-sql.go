package storage

import (
	"github.com/jinzhu/gorm"
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
	_ StorageStub = StorageGormSqlImpl{}
)

func NewStorageGormSql(dsn string) (*StorageGormSqlImpl, error) {
	conn, err := gorm.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	conn.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(GormToken{})
	return &StorageGormSqlImpl{conn}, nil
}

func (s StorageGormSqlImpl) Get(key string) (Token, error) {
	rt := new(GormToken)
	where := &GormToken{Token: Token{Token: []byte(key)}}
	if err := s.db.Where(where).First(rt).Error; err != nil {
		return Token{}, err
	}
	return rt.Token, nil
}

func (s StorageGormSqlImpl) Set(key string, value Token) error {
	s.Delete(key)
	return s.db.Create(&value).Error
}

func (s StorageGormSqlImpl) Delete(key string) error {
	where := &GormToken{Token: Token{Token: []byte(key)}}
	return s.db.Where(where).Delete(&GormToken{}).Error
}
