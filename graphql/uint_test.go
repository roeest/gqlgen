package graphql

import (
	"math"
	"testing"

	"github.com/goccy/go-json"

	"github.com/stretchr/testify/assert"
)

func TestUint(t *testing.T) {
	t.Run("marshal", func(t *testing.T) {
		assert.Equal(t, "123", m2s(MarshalUint(123)))
	})

	t.Run("unmarshal", func(t *testing.T) {
		assert.Equal(t, uint(123), mustUnmarshalUint(123))
		assert.Equal(t, uint(123), mustUnmarshalUint(int64(123)))
		assert.Equal(t, uint(123), mustUnmarshalUint(json.Number("123")))
		assert.Equal(t, uint(123), mustUnmarshalUint("123"))
	})

	t.Run("can't unmarshal negative numbers", func(t *testing.T) {
		_, err := UnmarshalUint(-123)
		assert.EqualError(t, err, "cannot convert negative numbers to uint")
	})
}

func mustUnmarshalUint(v any) uint {
	res, err := UnmarshalUint(v)
	if err != nil {
		panic(err)
	}
	return res
}

func TestUint32(t *testing.T) {
	t.Run("marshal", func(t *testing.T) {
		assert.Equal(t, "123", m2s(MarshalUint32(123)))
		assert.Equal(t, "4294967295", m2s(MarshalUint32(math.MaxUint32)))
	})

	t.Run("unmarshal", func(t *testing.T) {
		assert.Equal(t, uint32(123), mustUnmarshalUint32(123))
		assert.Equal(t, uint32(123), mustUnmarshalUint32(int64(123)))
		assert.Equal(t, uint32(123), mustUnmarshalUint32(json.Number("123")))
		assert.Equal(t, uint32(123), mustUnmarshalUint32("123"))
		assert.Equal(t, uint32(4294967295), mustUnmarshalUint32("4294967295"))
	})

	t.Run("can't unmarshal negative numbers", func(t *testing.T) {
		_, err := UnmarshalUint32(-123)
		assert.EqualError(t, err, "cannot convert negative numbers to uint32")
	})
}

func mustUnmarshalUint32(v any) uint32 {
	res, err := UnmarshalUint32(v)
	if err != nil {
		panic(err)
	}
	return res
}

func TestUint64(t *testing.T) {
	t.Run("marshal", func(t *testing.T) {
		assert.Equal(t, "123", m2s(MarshalUint64(123)))
	})

	t.Run("unmarshal", func(t *testing.T) {
		assert.Equal(t, uint64(123), mustUnmarshalUint64(123))
		assert.Equal(t, uint64(123), mustUnmarshalUint64(int64(123)))
		assert.Equal(t, uint64(123), mustUnmarshalUint64(json.Number("123")))
		assert.Equal(t, uint64(123), mustUnmarshalUint64("123"))
	})

	t.Run("can't unmarshal negative numbers", func(t *testing.T) {
		_, err := UnmarshalUint64(-123)
		assert.EqualError(t, err, "cannot convert negative numbers to uint64")
	})
}

func mustUnmarshalUint64(v any) uint64 {
	res, err := UnmarshalUint64(v)
	if err != nil {
		panic(err)
	}
	return res
}
