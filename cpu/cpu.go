package cpu

import (
	"bufio"
	"io"
	"os"
)

type CPU struct {
	code []Op
	tape []byte
	ir   Op   // instruction register
	ip   uint // instruction pointer
	tp   uint // tape pointer
	ok   bool // health indicator
	err  any
	in   *bufio.Reader
	out  *bufio.Writer
}

type Op func(*CPU)

func Default(code []Op) *CPU {
	return New(code, os.Stdin, os.Stdout)
}

func New(code []Op, in io.Reader, out io.Writer) *CPU {
	return &CPU{
		code: code,
		tape: make([]byte, initialTapeLength),
		ir:   NOP,
		ip:   0,
		tp:   initialTapeLength / 2,
		ok:   true,
		in:   bufio.NewReader(in),
		out:  bufio.NewWriter(out),
	}
}

func (c *CPU) Start() any {
	defer c.recover()
	for c.ok {
		c.fetch().execute()
	}
	return c.err
}

func (c *CPU) fetch() *CPU {
	if c.hasNoMoreOps() {
		c.ok = false
		c.ir = NOP
	} else {
		c.ir = c.code[c.ip]
		c.ip++
	}
	return c
}

func (c *CPU) hasNoMoreOps() bool {
	return c.ip >= uint(len(c.code))
}

func (c *CPU) execute() {
	c.ir(c)
}

func (c *CPU) recover() {
	c.err = recover()
}
