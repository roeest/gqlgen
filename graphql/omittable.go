package graphql

import (
	"encoding/json"

	fastjson "github.com/goccy/go-json"
)

// Omittable is a wrapper around a value that also stores whether it is set
// or not.
type Omittable[T any] struct {
	value T
	set   bool
}

var (
	_ json.Marshaler   = Omittable[struct{}]{}
	_ json.Unmarshaler = (*Omittable[struct{}])(nil)
)

func OmittableOf[T any](value T) Omittable[T] {
	return Omittable[T]{
		value: value,
		set:   true,
	}
}

func (o Omittable[T]) Value() T {
	if !o.set {
		var zero T
		return zero
	}
	return o.value
}

func (o Omittable[T]) ValueOK() (T, bool) {
	if !o.set {
		var zero T
		return zero, false
	}
	return o.value, true
}

func (o Omittable[T]) IsSet() bool {
	return o.set
}

func (o Omittable[T]) MarshalJSON() ([]byte, error) {
	if !o.set {
		return []byte("null"), nil
	}
	return fastjson.Marshal(o.value)
}

func (o *Omittable[T]) UnmarshalJSON(bytes []byte) error {
	err := fastjson.Unmarshal(bytes, &o.value)
	if err != nil {
		return err
	}
	o.set = true
	return nil
}
