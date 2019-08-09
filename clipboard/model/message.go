package model

import (
	"encoding/json"

	"gitlab.com/orbli/clipboard/util/storage"
)

type (
	Message struct {
		KeyString string
		Value     string
	}
)

var (
	_ storage.Value = &Message{}
)

func (m Message) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}

func (m Message) UnmarshalBinary(d []byte) error {
	return json.Unmarshal(d, m)
}

func (m Message) Key() string {
	return m.KeyString
}
