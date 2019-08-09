package storage

import "encoding"

type (
	Value interface {
		encoding.BinaryMarshaler
		encoding.BinaryUnmarshaler
		Key() string
	}
)
