package parse

import (
	"bufio"
	"io"
)

type stream struct {
	buf *bufio.Reader
}

func newBufferedStream(r io.Reader) *stream {
	return &stream{bufio.NewReader(r)}
}

func (s *stream) Next() (b byte, err error) {
	return s.buf.ReadByte()
}
