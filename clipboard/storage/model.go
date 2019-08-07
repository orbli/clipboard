package storage

import (
	"encoding"
	"encoding/json"
)

type (
	Message struct {
		Key   string
		Value string
	}
)

var (
	_ encoding.BinaryMarshaler   = Message{}
	_ encoding.BinaryUnmarshaler = &Message{}
)

func (m Message) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}

func (m *Message) UnmarshalBinary(d []byte) error {
	return json.Unmarshal(d, m)
}
