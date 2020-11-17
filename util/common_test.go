package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertByteToUint32(t *testing.T) {
	buf := []byte{0, 0, 1, 4}
	v := ConvertByteToUint32(buf, 0)
	assert.Equal(t, uint32(260), v)
}

func TestSliceContains(t *testing.T) {
	v := buildStringSlice("a", "b", "cd")
	assert.True(t, SliceContains(v, "cd"))
	assert.False(t, SliceContains(v, "d"))
}

func buildStringSlice(v ...interface{}) []interface{} {
	return v
}
