package util

import (
	"errors"
	"io"
	"sync"
)

// DualChannel: not concurrent safe
type DualChannel struct {
	buf        []byte
	readOffset int
}

func NewDualChannel() *DualChannel {
	return &DualChannel{buf: make([]byte, 0), readOffset: 0}
}

func (c *DualChannel) Write(data []byte) (int, error) {
	c.buf = append(c.buf, data...)
	return len(data), nil
}

func (c *DualChannel) Read(buf []byte) (int, error) {
	if c.readOffset >= len(c.buf) {
		return 0, io.EOF
	}
	n := min(len(buf), len(c.buf)-c.readOffset)
	for i := 0; i < n; i++ {
		buf[i] = c.buf[c.readOffset+i]
	}
	c.readOffset += n
	return n, nil
}

func (c *DualChannel) Reset() {
	c.buf = make([]byte, 0)
	c.readOffset = 0
}

const syncDualChannelBufLimit = 16 * 1024 // 16kb

type SyncDualChannel struct {
	buf        []byte
	readOffset int
	closed     bool
	mu         sync.Mutex
	sync       chan bool
}

func NewSyncDualChannel() *SyncDualChannel {
	return &SyncDualChannel{
		buf:        make([]byte, 0),
		sync:       make(chan bool, 1),
		readOffset: 0,
	}
}

func (c *SyncDualChannel) Write(data []byte) (int, error) {
	c.mu.Lock()
	if c.closed {
		c.mu.Unlock()
		return 0, errors.New("channel has been closed")
	}
	c.buf = append(c.buf, data...)
	c.mu.Unlock()

	// to prevent blocking routine, check channel size before writing
	if len(c.sync) == 0 {
		c.sync <- true
	}
	return len(data), nil
}

func (c *SyncDualChannel) Read(buf []byte) (int, error) {
	bufSz := len(buf)

	var readable int
	c.mu.Lock()
	for readable = len(c.buf) - c.readOffset; readable < bufSz; {
		if c.closed {
			c.mu.Unlock()
			return 0, io.EOF
		}
		c.mu.Unlock()

		<-c.sync
		c.mu.Lock()
		readable = len(c.buf) - c.readOffset
	}
	c.mu.Unlock()

	for i := 0; i < bufSz; i++ {
		buf[i] = c.buf[c.readOffset+i]
	}
	c.readOffset += bufSz

	if c.readOffset >= syncDualChannelBufLimit {
		// cut down buffer
		c.mu.Lock()
		c.buf = c.buf[c.readOffset:]
		c.readOffset = 0
		c.mu.Unlock()
	}

	return readable, nil
}

func (c *SyncDualChannel) Reset() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.buf = make([]byte, 0)
	c.readOffset = 0
}

func (c *SyncDualChannel) Close() {
	c.mu.Lock()
	c.closed = true
	c.mu.Unlock()

	if len(c.sync) == 0 {
		c.sync <- true
	}
}
