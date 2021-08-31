package gob

import (
	"bytes"
	"encoding/gob"

	"github.com/Sef1995/go-kit/encoding"
	"github.com/Sef1995/go-kit/types"
)

type encoder struct{}

// NewEncoder returns a new gob encoder.
func NewEncoder() encoding.Encoder {
	return &encoder{}
}

// Encode encodes the given value.
func (*encoder) Encode(obj interface{}) (*string, error) {
	var buffer bytes.Buffer

	if err := gob.NewEncoder(&buffer).Encode(obj); err != nil {
		return nil, err
	}

	return types.String(buffer.String()), nil
}

// Decode decodes the given value.
func (*encoder) Decode(encoded string, obj interface{}) error {
	var buffer bytes.Buffer

	buffer.WriteString(encoded)

	return gob.NewDecoder(&buffer).Decode(obj)
}
