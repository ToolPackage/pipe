package executor

import "io"

type DualChannel struct {
	buf        []byte
	readOffset int
}

func NewDualChannel() *DualChannel {
	return &DualChannel{buf: make([]byte, 0), readOffset: 0}
}

func (s *DualChannel) Write(data []byte) (int, error) {
	s.buf = append(s.buf, data...)
	return len(data), nil
}

func (s *DualChannel) Read(buf []byte) (int, error) {
	if s.readOffset >= len(s.buf) {
		return 0, io.EOF
	}
	n := min(len(buf), len(s.buf)-s.readOffset)
	for i := 0; i < n; i++ {
		buf[i] = s.buf[s.readOffset+i]
	}
	s.readOffset += n
	return n, nil
}

func (s *DualChannel) Reset() {
	s.buf = make([]byte, 0)
	s.readOffset = 0
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
