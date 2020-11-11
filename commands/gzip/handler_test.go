package gzip

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestGzip(t *testing.T) {
	data := "asd"
	buf := new(strings.Builder)
	_ = Compress(nil, bytes.NewReader([]byte(data)), buf)
	compressedData := []byte(buf.String())
	buf = new(strings.Builder)
	_ = Decompress(nil, bytes.NewReader(compressedData), buf)
	assert.Equal(t, data, buf.String())
}
