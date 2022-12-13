package cpu

import (
	"bytes"
	"os"
	"testing"
)

func TestCellOverflow(t *testing.T) {
	code := make([]Op, 0, 300)
	for i := 0; i < 300; i++ {
		code = append(code, INC)
	}
	vm := Default(code)
	if err := vm.Start(); err != nil {
		t.Fatalf("panic during execution: %v", err)
	}
	const expect = 44
	if cell := vm.tape[vm.tp]; cell != expect {
		t.Fatalf("failed to properly overflow cell value: wanted %d; got %d",
			expect, cell)
	}
}

func TestStdIn(t *testing.T) {
	code := []Op{IN}
	const letter = 'a'
	vm := New(code, bytes.NewReader([]byte{letter}), os.Stdout)
	if err := vm.Start(); err != nil {
		t.Fatalf("panic during execution: %v", err)
	}
	if cell := vm.tape[vm.tp]; cell != letter {
		t.Fatalf("failed to read from stdin: wanted %v; got %v", letter, cell)
	}
}

func TestStdOut(t *testing.T) {
	code := []Op{INC, INC, INC, INC, INC, OUT}
	buf := bytes.NewBuffer([]byte{})
	vm := New(code, os.Stdin, buf)
	if err := vm.Start(); err != nil {
		t.Fatalf("panic during execution: %v", err)
	}
	if buf.Len() != 1 {
		t.Fatalf("invalid buffer length: wanted %d; got %d", 1, buf.Len())
	}
	b, err := buf.ReadByte()
	if err != nil {
		t.Fatalf("failed to read from buf: %s", err)
	}
	if b != 5 {
		t.Fatalf("invalid byte written to buffer: wanted %d; got %d", 5, b)
	}
}
