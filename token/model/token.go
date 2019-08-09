package model

import (
	"encoding/json"
	"time"

	"gitlab.com/orbli/clipboard/util/storage"
)

type (
	Token struct {
		Token    []byte
		Secret   []byte
		Parent   string
		Data     []byte
		ExpireAt *time.Time
	}
)

var (
	_ storage.Value = &Token{}
)

func (m Token) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}

func (m Token) UnmarshalBinary(d []byte) error {
	return json.Unmarshal(d, m)
}

func (m Token) Key() string {
	return string(m.Token)
}
