package util

import (
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
	"time"
)

func TestSyncDualChannel(t *testing.T) {
	c := NewSyncDualChannel()
	go func() {
		data := make([]byte, syncDualChannelBufLimit-1)
		data[512] = 'A'
		_, err := c.Write(data)
		assert.Nil(t, err)
		time.Sleep(250 * time.Millisecond)

		data[512] += 1
		_, err = c.Write(data)
		assert.Nil(t, err)
		time.Sleep(250 * time.Millisecond)

		data[512] += 1
		_, err = c.Write(data)
		assert.Nil(t, err)
		time.Sleep(250 * time.Millisecond)

		c.Close()

		_, err = c.Write(data)
		assert.NotNil(t, err)
	}()

	buf := make([]byte, syncDualChannelBufLimit-1)
	for i := 0; i < 3; i++ {
		_, err := c.Read(buf)
		assert.Nil(t, err)
		assert.True(t, buf[512] == byte('A'+i))
	}
	_, err := c.Read(buf)
	assert.True(t, err == io.EOF)
	assert.True(t, c.closed)
	assert.True(t, syncDualChannelBufLimit-1 == len(c.buf))

	c.Reset()
}
