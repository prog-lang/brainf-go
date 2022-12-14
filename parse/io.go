package parse

import (
	"io"
	"os"

	"github.com/prog-lang/brainf-go/cpu"
	"github.com/sharpvik/fungi"
)

func FromFile(name string) (ops []cpu.Op, err error) {
	file, err := os.Open(name)
	if err != nil {
		return
	}
	defer file.Close()
	return FromReader(file)
}

func FromReader(r io.Reader) ([]cpu.Op, error) {
	s := newBufferedStream(r)
	p := newParser()
	if err := fungi.ForEach(p.feed)(s); err != nil {
		return nil, err
	}
	return p.code()
}
