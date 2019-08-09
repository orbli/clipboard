package model

import (
	"encoding/json"
	"strconv"

	"gitlab.com/orbli/clipboard/util/storage"
)

type (
	User struct {
		Id       uint64
		Name     string
		Metadata []byte
	}
)

var (
	_ storage.Value = &User{}
)

func (m User) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}

func (m User) UnmarshalBinary(d []byte) error {
	return json.Unmarshal(d, m)
}

func (m User) Key() string {
	return strconv.FormatUint(m.Id, 10)
}
