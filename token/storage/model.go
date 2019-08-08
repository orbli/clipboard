package storage

import (
	"encoding"
	"encoding/json"
	"time"
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
	_ encoding.BinaryMarshaler   = Token{}
	_ encoding.BinaryUnmarshaler = &Token{}
)

func (m Token) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}

func (m *Token) UnmarshalBinary(d []byte) error {
	return json.Unmarshal(d, m)
}
